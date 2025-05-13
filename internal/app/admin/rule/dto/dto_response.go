package rule

import model "Synapse/internal/app/admin/rule/model"

// AdminRuleDTO representa uma regra (papel) retornada na listagem
type AdminRuleDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// AdminRuleListDTO representa uma resposta paginada de regras
type AdminRuleListDTO struct {
	Total  int64          `json:"total"`
	Limit  int            `json:"limit"`
	Offset int            `json:"offset"`
	Data   []AdminRuleDTO `json:"data"`
}

// RulePermissionDTO representa uma permissão no formato namespace
type RulePermissionDTO struct {
	Namespace string `json:"namespace"`
}

// FromModel converte um model.AdminRule para um DTO
func FromModel(rule model.AdminRule) AdminRuleDTO {
	return AdminRuleDTO{
		ID:   rule.ID,
		Name: rule.Name,
	}
}

// FromModelList converte uma lista de model.AdminRule para []AdminRuleDTO
func FromModelList(rules []model.AdminRule) []AdminRuleDTO {
	dtoList := make([]AdminRuleDTO, 0, len(rules))
	for _, rule := range rules {
		dtoList = append(dtoList, FromModel(rule))
	}
	return dtoList
}

// FromNamespaceList converte uma lista de strings (namespaces) em []RulePermissionDTO
func FromNamespaceList(namespaces []string) []RulePermissionDTO {
	dtoList := make([]RulePermissionDTO, 0, len(namespaces))
	for _, ns := range namespaces {
		dtoList = append(dtoList, RulePermissionDTO{Namespace: ns})
	}
	return dtoList
}
