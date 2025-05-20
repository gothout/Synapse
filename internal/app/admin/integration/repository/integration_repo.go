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
