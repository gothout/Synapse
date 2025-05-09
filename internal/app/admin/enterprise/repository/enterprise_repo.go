package enterprise

import (
	enterprise "Synapse/internal/app/admin/enterprise/model"

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
	//
	return nil, nil
}
