package enterprise

import (
	enterprise "Synapse/internal/app/admin/enterprise/model"
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

// Funcao para criar empresa.
func (r *repository) Create(enterprise *enterprise.AdminEnterprise) (*enterprise.AdminEnterprise, error) {
	ctx := context.Background()

	// Verifica se não existe empresa com o mesmo CNPJ.
	checkQuery := `SELECT id FROM admin_enterprise WHERE cnpj = $1`
	var idExistente int64
	err := r.db.QueryRow(ctx, checkQuery, enterprise.Cnpj).Scan(&idExistente)
	if err != nil && err != pgx.ErrNoRows {
		return nil, fmt.Errorf("erro ao verificar existência de empresa: %w", err)
	}
	if err == nil {
		return nil, fmt.Errorf("já existe uma empresa com este CNPJ")
	}

	enterprise.Cnpj = strings.TrimSpace(enterprise.Cnpj)

	// Inserir empresa
	insertQuery := `INSERT INTO admin_enterprise (nome, cnpj, created_at) VALUES ($1, $2, $3) RETURNING id`
	err = r.db.QueryRow(ctx, insertQuery, enterprise.Nome, enterprise.Cnpj, enterprise.CreatedAt).Scan(&enterprise.ID)
	if err != nil {
		return nil, fmt.Errorf("erro ao inserir empresa: %w", err)
	}
	return enterprise, nil
}
