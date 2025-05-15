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

// Funcao para buscar empresa por CNPJ
func (r *repository) ReadByCNPJ(cnpj string) (*enterprise.AdminEnterprise, error) {
	ctx := context.Background()

	query := `SELECT id, nome, cnpj, created_at,update_at FROM admin_enterprise WHERE cnpj = $1`
	var enterprise enterprise.AdminEnterprise

	err := r.db.QueryRow(ctx, query, cnpj).Scan(&enterprise.ID, &enterprise.Nome, &enterprise.Cnpj, &enterprise.CreatedAt, &enterprise.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("empresa com CNPJ não encontrada")
		}
		return nil, fmt.Errorf("erro ao buscar empresa por CNPJ")
	}

	return &enterprise, nil
}

// Busca empresa por NOME
func (r *repository) ReadByNome(nome string) (*enterprise.AdminEnterprise, error) {
	ctx := context.Background()

	query := `SELECT id, nome, cnpj, created_at, update_at FROM admin_enterprise WHERE nome = $1`
	var enterprise enterprise.AdminEnterprise

	err := r.db.QueryRow(ctx, query, nome).Scan(&enterprise.ID, &enterprise.Nome, &enterprise.Cnpj, &enterprise.CreatedAt, &enterprise.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("empresa com NOME não encontrada")
		}
		return nil, fmt.Errorf("erro ao buscar empresa por NOME")
	}

	return &enterprise, nil
}

// Busca empresa por ID
func (r *repository) ReadByID(id int64) (*enterprise.AdminEnterprise, error) {
	ctx := context.Background()
	query := `SELECT id, nome, cnpj, created_at,update_at FROM admin_enterprise WHERE id = $1`
	var enterprise enterprise.AdminEnterprise

	err := r.db.QueryRow(ctx, query, id).Scan(&enterprise.ID, &enterprise.Nome, &enterprise.Cnpj, &enterprise.CreatedAt, &enterprise.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {

			return nil, fmt.Errorf("empresa com ID não encontrada")
		}
		return nil, fmt.Errorf("erro ao buscar empresa por ID")
	}

	return &enterprise, nil
}

// Lista empreasas por page
func (r *repository) ReadAll(page int) ([]enterprise.AdminEnterprise, error) {
	ctx := context.Background()

	var rows pgx.Rows
	var err error

	// Se page == -1, listar tudo
	if page == -1 {
		rows, err = r.db.Query(ctx, `SELECT id, nome, cnpj, created_at, update_at FROM admin_enterprise ORDER BY id`)
	} else if page >= 1 {
		limit := 10
		offset := (page - 1) * limit
		query := `SELECT id, nome, cnpj, created_at, update_at FROM admin_enterprise ORDER BY id LIMIT $1 OFFSET $2`
		rows, err = r.db.Query(ctx, query, limit, offset)
	} else {
		return nil, fmt.Errorf("número de página inválido")
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao buscar empresas: %w", err)
	}
	defer rows.Close()

	var empresas []enterprise.AdminEnterprise
	for rows.Next() {
		var e enterprise.AdminEnterprise
		if err := rows.Scan(&e.ID, &e.Nome, &e.Cnpj, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, fmt.Errorf("erro ao fazer scan: %w", err)
		}
		empresas = append(empresas, e)
	}

	return empresas, nil
}

// Atualiza empresa baseado no CNPJ
func (r *repository) UpdateByCNPJ(cnpj, newCNPJ string, updated *enterprise.AdminEnterprise) (*enterprise.AdminEnterprise, error) {
	ctx := context.Background()
	updated.Nome = strings.TrimSpace(updated.Nome)
	newCNPJ = strings.TrimSpace(newCNPJ)

	// Verifica se a empresa existe
	checkExistence := `SELECT id FROM admin_enterprise WHERE cnpj = $1`
	var existingID int64
	err := r.db.QueryRow(ctx, checkExistence, cnpj).Scan(&existingID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("empresa com CNPJ fornecido não encontrada")
		}
		return nil, fmt.Errorf("erro ao verificar existência da empresa: %w", err)
	}

	// Lógica de atualização
	switch {
	case updated.Nome != "" && newCNPJ == "":
		// Atualiza apenas nome
		query := `UPDATE admin_enterprise SET nome = $1, update_at = $2 WHERE cnpj = $3`
		_, err := r.db.Exec(ctx, query, updated.Nome, updated.UpdatedAt, cnpj)
		if err != nil {
			return nil, fmt.Errorf("erro ao atualizar nome: %w", err)
		}
		updated.Cnpj = cnpj

	case updated.Nome == "" && newCNPJ != "":
		if newCNPJ != cnpj {
			if err := checkDuplicateCNPJ(r, newCNPJ); err != nil {
				return nil, err
			}
		}
		query := `UPDATE admin_enterprise SET cnpj = $1, update_at = $2 WHERE cnpj = $3`
		_, err := r.db.Exec(ctx, query, newCNPJ, updated.UpdatedAt, cnpj)
		if err != nil {
			return nil, fmt.Errorf("erro ao atualizar CNPJ: %w", err)
		}
		updated.Cnpj = newCNPJ

	case updated.Nome != "" && newCNPJ != "":
		if newCNPJ != cnpj {
			if err := checkDuplicateCNPJ(r, newCNPJ); err != nil {
				return nil, err
			}
		}
		query := `UPDATE admin_enterprise SET nome = $1, cnpj = $2, update_at = $3 WHERE cnpj = $4`
		_, err := r.db.Exec(ctx, query, updated.Nome, newCNPJ, updated.UpdatedAt, cnpj)
		if err != nil {
			return nil, fmt.Errorf("erro ao atualizar nome e CNPJ: %w", err)
		}
		updated.Cnpj = newCNPJ

	default:
		return nil, fmt.Errorf("nenhum dado para atualizar foi fornecido")
	}

	return updated, nil
}

// DeleteByCNPJ remove uma empresa com base no CNPJ fornecido.
// Retorna o CNPJ excluído, ou erro se não encontrada ou falha no banco.
func (r *repository) DeleteByCNPJ(cnpj string) (string, error) {
	ctx := context.Background()

	// Primeiro verifica se a empresa existe
	var existingID int64
	checkQuery := `SELECT id FROM admin_enterprise WHERE cnpj = $1`
	err := r.db.QueryRow(ctx, checkQuery, cnpj).Scan(&existingID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("empresa com CNPJ fornecido não encontrada")
		}
		return "", fmt.Errorf("erro ao verificar existência da empresa: %w", err)
	}

	// Executa a exclusão
	deleteQuery := `DELETE FROM admin_enterprise WHERE cnpj = $1`
	_, err = r.db.Exec(ctx, deleteQuery, cnpj)
	if err != nil {
		return "", fmt.Errorf("erro ao excluir empresa: %w", err)
	}

	return cnpj, nil
}

// Helper
func checkDuplicateCNPJ(r *repository, cnpj string) error {
	checkQuery := `SELECT id FROM admin_enterprise WHERE cnpj = $1`
	var existingID int64
	err := r.db.QueryRow(context.Background(), checkQuery, cnpj).Scan(&existingID)
	if err != nil && err != pgx.ErrNoRows {
		return fmt.Errorf("erro ao verificar CNPJ duplicado: %w", err)
	}
	if err == nil {
		return fmt.Errorf("já existe uma empresa com este novo CNPJ")
	}
	return nil
}
