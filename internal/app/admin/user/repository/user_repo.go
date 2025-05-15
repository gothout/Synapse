package user

import (
	user "Synapse/internal/app/admin/user/model"
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db}
}

// Create insere um novo usuário no banco de dados.
func (r *repository) Create(user *user.User) (*user.User, error) {
	ctx := context.Background()

	// Verifica se já existe usuário com o mesmo e-mail
	checkQuery := `SELECT id FROM admin_user WHERE email = $1`
	var idExistente int64
	err := r.db.QueryRow(ctx, checkQuery, user.Email).Scan(&idExistente)
	if err != nil && err != pgx.ErrNoRows {
		return nil, fmt.Errorf("erro ao verificar existência de usuário: %w", err)
	}
	if err == nil {
		return nil, fmt.Errorf("já existe um usuário com este e-mail")
	}

	// Remove espaços em branco
	user.Email = strings.TrimSpace(user.Email)
	user.Nome = strings.TrimSpace(user.Nome)
	user.Numero = strings.TrimSpace(user.Numero)
	user.Token = strings.TrimSpace(user.Token)

	// Insere o usuário
	insertQuery := `
		INSERT INTO admin_user (nome, email, senha, numero, token, rule_id, enterprise_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id
	`

	err = r.db.QueryRow(
		ctx, insertQuery,
		user.Nome,
		user.Email,
		user.Senha,
		user.Numero,
		user.Token,
		user.RuleID,
		user.EnterpriseID,
	).Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("erro ao inserir usuário: %w", err)
	}

	return user, nil
}
