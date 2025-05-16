package user

import (
	"Synapse/internal/app/admin/pkg/security"
	user "Synapse/internal/app/admin/user/model"
	"context"
	"fmt"
	"strings"
	"time"

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

// ReadByID retorna um usuário com base no ID.
func (r *repository) ReadByID(userID int64) (*user.User, error) {
	ctx := context.Background()

	query := `
		SELECT id, nome, email, senha, numero, token, rule_id, enterprise_id, created_at, updated_at
		FROM admin_user
		WHERE id = $1
	`

	var u user.User
	err := r.db.QueryRow(ctx, query, userID).Scan(
		&u.ID,
		&u.Nome,
		&u.Email,
		&u.Senha,
		&u.Numero,
		&u.Token,
		&u.RuleID,
		&u.EnterpriseID,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("usuário com ID %d não encontrado", userID)
		}
		return nil, fmt.Errorf("erro ao buscar usuário por ID: %w", err)
	}

	return &u, nil
}

// Atualiza usuario por ID por ID
func (r *repository) UpdateUserByID(userID int64, updated *user.User) (*user.User, error) {
	ctx := context.Background()

	// Verifica se o usuário existe
	queryCheck := `SELECT id FROM admin_user WHERE id = $1`
	var existingID int64
	err := r.db.QueryRow(ctx, queryCheck, userID).Scan(&existingID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("usuário com ID %d não encontrado", userID)
		}
		return nil, fmt.Errorf("erro ao verificar existência do usuário: %w", err)
	}

	// Se a senha vier vazia, busca a atual para manter
	if updated.Senha == "" {
		querySenha := `SELECT senha FROM admin_user WHERE id = $1`
		err := r.db.QueryRow(ctx, querySenha, userID).Scan(&updated.Senha)
		if err != nil {
			return nil, fmt.Errorf("erro ao obter senha atual: %w", err)
		}
	}

	// Atualiza os campos, incluindo enterprise_id
	queryUpdate := `
		UPDATE admin_user 
		SET nome = $1, email = $2, senha = $3, numero = $4, rule_id = $5, enterprise_id = $6, updated_at = CURRENT_TIMESTAMP
		WHERE id = $7
		RETURNING id, nome, email, senha, numero, token, rule_id, enterprise_id, created_at, updated_at
	`

	row := r.db.QueryRow(ctx, queryUpdate,
		updated.Nome,
		updated.Email,
		updated.Senha,
		updated.Numero,
		updated.RuleID,
		updated.EnterpriseID,
		userID,
	)

	var updatedUser user.User
	err = row.Scan(
		&updatedUser.ID,
		&updatedUser.Nome,
		&updatedUser.Email,
		&updatedUser.Senha,
		&updatedUser.Numero,
		&updatedUser.Token,
		&updatedUser.RuleID,
		&updatedUser.EnterpriseID,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	return &updatedUser, nil
}

// DeleteUserByID remove um usuário com base no ID.
func (r *repository) DeleteUserByID(userID int64) error {
	ctx := context.Background()

	query := `DELETE FROM admin_user WHERE id = $1`

	commandTag, err := r.db.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("erro ao deletar usuário")
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("erro ao deletar usuário")
		//return fmt.Errorf("nenhum usuário encontrado")
	}

	return nil
}

// Autenticação
func (r *repository) ValidateCredentials(ctx context.Context, email, senha string) (*user.User, error) {
	var user user.User

	query := `SELECT id, nome, email, senha, numero, rule_id, enterprise_id FROM admin_user WHERE email = $1`
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Nome,
		&user.Email,
		&user.Senha,
		&user.Numero,
		&user.RuleID,
		&user.EnterpriseID,
	)
	if err != nil {
		return nil, err
	}

	if !security.CheckPasswordHash(senha, user.Senha) {
		return nil, fmt.Errorf("senha inválida")
	}

	return &user, nil
}

// Salvar token
func (r *repository) SaveToken(ctx context.Context, userID int64, token string, expireAt time.Time) error {
	var count int

	// Verifica se já existe um token válido
	queryCheck := `
		SELECT COUNT(*) FROM admin_token
		WHERE user_id = $1 AND expires_at > CURRENT_TIMESTAMP
	`
	err := r.db.QueryRow(ctx, queryCheck, userID).Scan(&count)
	if err != nil {
		return fmt.Errorf("erro ao verificar token existente: %w", err)
	}

	// Se já existe token válido, não cria novo
	if count > 0 {
		return nil
	}

	// Insere novo token com expireAt vindo do JWT
	queryInsert := `
		INSERT INTO admin_token (token, user_id, created_at, expires_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP, $3)
	`
	_, err = r.db.Exec(ctx, queryInsert, token, userID, expireAt.UTC()) // força UTC
	if err != nil {
		return fmt.Errorf("erro ao inserir novo token: %w", err)
	}

	return nil
}

// Valida um token se esta valido ou expirado.
func (r *repository) GetValidToken(ctx context.Context, userID int64) (string, error) {
	var token string
	query := `SELECT token FROM admin_token WHERE user_id = $1 AND expires_at > CURRENT_TIMESTAMP LIMIT 1`
	err := r.db.QueryRow(ctx, query, userID).Scan(&token)
	if err != nil {
		return "", nil // Nenhum token válido
	}
	return token, nil
}
