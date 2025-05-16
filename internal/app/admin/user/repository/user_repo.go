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

// ReadAllUser retorna todos os usuários da empresa com paginação.
func (r *repository) ReadAllUser(enterpriseID int64, page int64) (*[]user.User, error) {
	ctx := context.Background()
	var rows pgx.Rows
	var err error

	const limit = 10
	offset := (page - 1) * limit

	query := `
		SELECT id, nome, email, numero, token, rule_id, enterprise_id, created_at, updated_at
		FROM admin_user
		WHERE enterprise_id = $1
		ORDER BY id
		LIMIT $2 OFFSET $3
	`

	rows, err = r.db.Query(ctx, query, enterpriseID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuários: %w", err)
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		err := rows.Scan(&u.ID, &u.Nome, &u.Email, &u.Numero, &u.Token, &u.RuleID, &u.EnterpriseID, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear usuário: %w", err)
		}
		users = append(users, u)
	}

	return &users, nil
}

// ReadByEmail retorna um usuário com base no e-mail.
func (r *repository) ReadByEmail(email string) (*user.User, error) {
	ctx := context.Background()

	query := `
		SELECT id, nome, email, numero, token, rule_id, enterprise_id, created_at, updated_at
		FROM admin_user
		WHERE email = $1
	`

	var u user.User
	err := r.db.QueryRow(ctx, query, email).Scan(
		&u.ID,
		&u.Nome,
		&u.Email,
		&u.Numero,
		&u.Token,
		&u.RuleID,
		&u.EnterpriseID,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("usuário com e-mail %s não encontrado", email)
		}
		return nil, fmt.Errorf("erro ao buscar usuário por e-mail: %w", err)
	}

	return &u, nil
}
