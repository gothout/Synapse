package rule

import (
	"context"
	"fmt"

	rules "Synapse/internal/app/admin/rule/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

// repository implementa a interface Repository e mantém a conexão com o banco
type repository struct {
	db *pgxpool.Pool
}

// NewRuleRepository retorna uma nova instância do repositório de regras
func NewRuleRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

// FindAllRules busca todas as regras (papéis) com paginação
// Ex: SELECT * FROM admin_rule LIMIT 10 OFFSET 0
func (r *repository) FindAllRules(limit, offset int) ([]rules.AdminRule, error) {
	query := `
		SELECT id, name
		FROM admin_rule
		ORDER BY id
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar regras paginadas: %w", err)
	}
	defer rows.Close()

	var rulesList []rules.AdminRule
	for rows.Next() {
		var rule rules.AdminRule
		if err := rows.Scan(&rule.ID, &rule.Name); err != nil {
			return nil, fmt.Errorf("erro ao escanear regra: %w", err)
		}
		rulesList = append(rulesList, rule)
	}

	return rulesList, nil
}

// FindPermissionsByRuleID retorna uma lista de permissões completas (namespace) associadas a uma regra
// Ex: admin.enterprise.create, admin.user.read
func (r *repository) FindPermissionsByRuleID(ruleID int64) ([]string, error) {
	query := `
		SELECT am.name || '.' || ap.action AS namespace
		FROM admin_rule_permission arp
		JOIN admin_permission ap ON arp.permission_id = ap.id
		JOIN admin_module am ON ap.module_id = am.id
		WHERE arp.rule_id = $1
		ORDER BY am.name, ap.action
	`

	rows, err := r.db.Query(context.Background(), query, ruleID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar permissões da regra: %w", err)
	}
	defer rows.Close()

	var permissions []string
	for rows.Next() {
		var ns string
		if err := rows.Scan(&ns); err != nil {
			return nil, fmt.Errorf("erro ao escanear namespace: %w", err)
		}
		permissions = append(permissions, ns)
	}

	return permissions, nil
}

// FindRuleByID busca uma única regra (papel) pelo ID
// Ex: SELECT id, name FROM admin_rule WHERE id = $1
func (r *repository) FindRuleByID(ruleID int64) (*rules.AdminRule, error) {
	query := `SELECT id, name FROM admin_rule WHERE id = $1`

	var rule rules.AdminRule
	err := r.db.QueryRow(context.Background(), query, ruleID).Scan(&rule.ID, &rule.Name)
	if err != nil {
		return nil, fmt.Errorf("regra não encontrada: %w", err)
	}

	return &rule, nil
}
