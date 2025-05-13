package rule

import rules "Synapse/internal/app/admin/rule/model"

type Repository interface {
	// Retorna todas as regras paginadas
	FindAllRules(limit, offset int) ([]rules.AdminRule, error)

	// Retorna permissões de uma regra como: admin.enterprise.create
	FindPermissionsByRuleID(ruleID int64) ([]string, error)

	// Retorna uma regra específica
	FindRuleByID(ruleID int64) (*rules.AdminRule, error)
}
