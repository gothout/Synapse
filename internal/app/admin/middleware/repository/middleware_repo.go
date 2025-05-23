package middleware

import (
	model "Synapse/internal/app/admin/middleware/model"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type middlewareRepository struct {
	db *pgxpool.Pool
}

func NewMiddlewareRepository(db *pgxpool.Pool) MiddlewareRepository {
	return &middlewareRepository{db}
}

// GetUserIDByToken retorna o userID associado a um token válido e não expirado
func (r *middlewareRepository) GetUserIDByToken(ctx context.Context, token string) (int64, error) {
	var userID int64
	query := `
        SELECT user_id 
        FROM admin_token 
        WHERE token = $1 AND (expires_at IS NULL OR expires_at > NOW())
        LIMIT 1
    `
	err := r.db.QueryRow(ctx, query, token).Scan(&userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, errors.New("token inválido ou expirado")
		}
		return 0, err
	}

	return userID, nil
}

// CheckPermission verifica se o usuário possui permissão para um módulo + ação
func (r *middlewareRepository) CheckPermission(ctx context.Context, userID int64, module, action string) (bool, error) {
	query := `
        SELECT 1
        FROM admin_user u
        JOIN admin_rule_permission rp ON rp.rule_id = u.rule_id
        JOIN admin_permission p ON p.id = rp.permission_id
        JOIN admin_module m ON m.id = p.module_id
        WHERE u.id = $1
          AND m.name = $2
          AND p.action = $3
        LIMIT 1
    `
	var exists int
	err := r.db.QueryRow(ctx, query, userID, module, action).Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// FindIntegrationByToken retorna os dados da integração + permissões associadas, com base no token
func (r *middlewareRepository) FindIntegrationByToken(ctx context.Context, token string) (*model.IntegrationWithPermissions, error) {
	query := `
		SELECT i.id, i.nome, t.token, m.nome as permissao
		FROM admin_integracao_tokens t
		JOIN admin_integracoes i ON i.id = t.integracao_id
		JOIN admin_integracao_marcas m ON m.id = i.marca_id
		WHERE t.token = $1
	`

	rows, err := r.db.Query(ctx, query, token)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var integration *model.IntegrationWithPermissions
	perms := []string{}

	for rows.Next() {
		var id, nome, tkn, permissao string
		if err := rows.Scan(&id, &nome, &tkn, &permissao); err != nil {
			return nil, err
		}
		if integration == nil {
			integration = &model.IntegrationWithPermissions{
				ID:         id,
				Nome:       nome,
				Token:      tkn,
				Permissoes: []string{},
			}
		}
		if permissao != "" {
			perms = append(perms, permissao)
		}
	}

	if integration != nil {
		integration.Permissoes = perms
	}

	return integration, nil
}

// CheckEnterpriseToken valida se o token fornecido está associado a um usuário e se esse usuário tem permissão
// para a integração solicitada, além de verificar se a empresa desse usuário também está associada à integração.
func (r *middlewareRepository) CheckEnterpriseToken(ctx context.Context, token string, integrationID string) (*model.IntegrationWithPermissions, error) {
	query := `
		SELECT 
			i.id, 
			i.nome,
			t.token,
			u.enterprise_id,
			u.id as user_id
		FROM admin_integracao_tokens t
		JOIN admin_user u ON u.id = t.user_id
		JOIN admin_integracoes i ON i.id = t.integracao_id
		JOIN admin_integracao_user iu ON iu.user_id = u.id AND iu.integracao_id = i.id
		JOIN admin_integracao_enterprise ie ON ie.enterprise_id = u.enterprise_id AND ie.integracao_id = i.id
		WHERE t.token = $1 AND i.id = $2
		LIMIT 1
	`

	var integration model.IntegrationWithPermissions
	var enterpriseID, userID int64

	err := r.db.QueryRow(ctx, query, token, integrationID).Scan(
		&integration.ID,
		&integration.Nome,
		&integration.Token,
		&enterpriseID,
		&userID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("token inválido ou sem permissão para esta integração")
		}
		return nil, err
	}

	integration.EnterpriseID = enterpriseID
	integration.UserID = userID
	return &integration, nil
}
