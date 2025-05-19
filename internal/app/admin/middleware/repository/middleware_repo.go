package middleware

import (
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

// Retorna o userID associado ao token, se for válido e não expirado
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

// Verifica se o user possui permissão para um módulo + ação
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
