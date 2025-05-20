package integration

import (
	model "Synapse/internal/app/admin/integration/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// repository implementa a interface Repository e mantém a conexão com o banco
type repository struct {
	db *pgxpool.Pool
}

// NewRuleRepository retorna uma nova instância do repositório de regras
func NewIntegrationRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

// GetAllIntegrations retorna todas as integrações disponíveis
func (r *repository) GetAllIntegrations() ([]model.Integration, error) {
	query := `
		SELECT id, nome, marca_id
		FROM admin_integracoes
		ORDER BY id;
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar integrações: %w", err)
	}
	defer rows.Close()

	var list []model.Integration
	for rows.Next() {
		var integ model.Integration
		if err := rows.Scan(&integ.Id, &integ.Nome, &integ.MarcaId); err != nil {
			return nil, fmt.Errorf("erro ao escanear integração: %w", err)
		}
		list = append(list, integ)
	}

	return list, nil
}

func (r *repository) GetAllMarcas() ([]model.Marca, error) {
	query := `SELECT id, nome FROM admin_integracao_marcas ORDER BY id`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar marcas de integração: %w", err)
	}
	defer rows.Close()

	var list []model.Marca
	for rows.Next() {
		var m model.Marca
		if err := rows.Scan(&m.Id, &m.Name); err != nil {
			return nil, fmt.Errorf("erro ao escanear marca: %w", err)
		}
		list = append(list, m)
	}

	return list, nil
}

// Busca detalhes de X Integracoes para determinada empresa
func (r *repository) GetIntegracoesDetalhadasByMarcaID(marcaID int64) ([]model.IntegracaoDetalhada, error) {
	query := `
		SELECT 
			i.id AS id_integracao,
			m.id AS id_marca,
			i.nome AS nome_integracao,
			m.nome AS nome_marca
		FROM admin_integracoes i
		INNER JOIN admin_integracao_marcas m ON m.id = i.marca_id
		WHERE i.marca_id = $1
		ORDER BY i.id
	`

	rows, err := r.db.Query(context.Background(), query, marcaID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar integrações detalhadas da marca %d: %w", marcaID, err)
	}
	defer rows.Close()

	var list []model.IntegracaoDetalhada
	for rows.Next() {
		var item model.IntegracaoDetalhada
		if err := rows.Scan(
			&item.IdIntegracao,
			&item.IdMarca,
			&item.NomeIntegracao,
			&item.NomeMarca,
		); err != nil {
			return nil, fmt.Errorf("erro ao escanear integração detalhada: %w", err)
		}
		list = append(list, item)
	}

	return list, nil
}

// Vincular empresa a integração
func (r *repository) CreateIntegracaoEnterprise(data model.IntegracaoEnterprise) error {
	query := `
		INSERT INTO admin_integracao_enterprise (enterprise_id, integracao_id)
		VALUES ($1, $2)
		ON CONFLICT (enterprise_id, integracao_id) DO NOTHING
	`

	_, err := r.db.Exec(context.Background(), query, data.EnterpriseId, data.IntegracaoId)
	if err != nil {
		return fmt.Errorf("erro ao vincular integração à empresa")
	}

	return nil
}

func (r *repository) GetIntegracaoByID(ctx context.Context, id int64) (*model.Integration, error) {
	query := `
		SELECT id, nome, marca_id
		FROM admin_integracoes
		WHERE id = $1
	`

	var integ model.Integration
	err := r.db.QueryRow(ctx, query, id).Scan(&integ.Id, &integ.Nome, &integ.MarcaId)
	if err != nil {
		return nil, fmt.Errorf("integração com ID %d não encontrada: %w", id, err)
	}

	return &integ, nil
}

// Retorna detalhe de integrações para X empresa
func (r *repository) GetIntegracoesByEnterpriseID(enterpriseID int64) ([]model.IntegracaoEmpresaDetalhada, error) {
	query := `
		SELECT 
			i.id AS integracao_id,
			i.nome AS integracao,
			m.nome AS marca
		FROM admin_integracao_enterprise ie
		JOIN admin_integracoes i ON ie.integracao_id = i.id
		JOIN admin_integracao_marcas m ON i.marca_id = m.id
		WHERE ie.enterprise_id = $1
		ORDER BY i.id
	`

	rows, err := r.db.Query(context.Background(), query, enterpriseID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar integrações da empresa %d: %w", enterpriseID, err)
	}
	defer rows.Close()

	var result []model.IntegracaoEmpresaDetalhada
	for rows.Next() {
		var item model.IntegracaoEmpresaDetalhada
		if err := rows.Scan(&item.IntegracaoID, &item.Integracao, &item.Marca); err != nil {
			return nil, fmt.Errorf("erro ao escanear integração: %w", err)
		}
		result = append(result, item)
	}

	return result, nil
}

// remover a integração de uma empresa específica
func (r *repository) DeleteIntegracaoFromEnterprise(enterpriseID, integracaoID int64) error {
	query := `
		DELETE FROM admin_integracao_enterprise
		WHERE enterprise_id = $1 AND integracao_id = $2
	`

	cmdTag, err := r.db.Exec(context.Background(), query, enterpriseID, integracaoID)
	if err != nil {
		return fmt.Errorf("erro ao remover integração da empresa")
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("nenhuma relação encontrada para remover")
	}

	return nil
}

// Vincular usuario a integração
func (r *repository) CreateIntegracaoUser(data model.IntegracaoUser) error {
	query := `
		INSERT INTO admin_integracao_user (user_id, integracao_id)
		VALUES ($1, $2)
		ON CONFLICT (user_id, integracao_id) DO NOTHING
	`

	_, err := r.db.Exec(context.Background(), query, data.UserID, data.IntegracaoID)
	if err != nil {
		return fmt.Errorf("erro ao vincular usuário à integração: %w", err)
	}
	return nil
}

// Salvar token para acesso a integração
func (r *repository) SaveIntegracaoToken(userID, integracaoID int64, token string) error {
	ctx := context.Background()

	// Remove qualquer token antigo
	delQuery := `
		DELETE FROM admin_integracao_tokens
		WHERE user_id = $1 AND integracao_id = $2
	`
	if _, err := r.db.Exec(ctx, delQuery, userID, integracaoID); err != nil {
		return fmt.Errorf("erro ao remover token antigo: %w", err)
	}

	// Insere novo token
	insertQuery := `
		INSERT INTO admin_integracao_tokens (token, user_id, integracao_id)
		VALUES ($1, $2, $3)
	`
	if _, err := r.db.Exec(ctx, insertQuery, token, userID, integracaoID); err != nil {
		return fmt.Errorf("erro ao salvar novo token: %w", err)
	}

	return nil
}

// Checar o usuario pela integração
func (r *repository) CheckUserHasIntegracao(userID, integracaoID int64) (bool, error) {
	query := `
		SELECT COUNT(1)
		FROM admin_integracao_user
		WHERE user_id = $1 AND integracao_id = $2
	`
	var count int
	err := r.db.QueryRow(context.Background(), query, userID, integracaoID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("erro ao verificar vínculo do usuário à integração: %w", err)
	}
	return count > 0, nil
}

// Busca  os dados de uma integração que o usuário possui
func (r *repository) GetIntegracoesByUserID(userID int64) ([]model.IntegracaoUsuarioDetalhada, error) {
	query := `
		SELECT ai.id, ai.nome, aim.nome AS marca
		FROM admin_integracao_user aiu
		JOIN admin_integracoes ai ON aiu.integracao_id = ai.id
		JOIN admin_integracao_marcas aim ON ai.marca_id = aim.id
		WHERE aiu.user_id = $1
		ORDER BY ai.id
	`

	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar permissões de integração do usuário: %w", err)
	}
	defer rows.Close()

	var result []model.IntegracaoUsuarioDetalhada
	for rows.Next() {
		var i model.IntegracaoUsuarioDetalhada
		if err := rows.Scan(&i.ID, &i.Nome, &i.Marca); err != nil {
			return nil, fmt.Errorf("erro ao escanear resultado: %w", err)
		}
		result = append(result, i)
	}

	return result, nil
}
