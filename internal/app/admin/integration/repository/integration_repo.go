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
