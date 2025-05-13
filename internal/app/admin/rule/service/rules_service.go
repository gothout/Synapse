package rule

import (
	model "Synapse/internal/app/admin/rule/model"
	repository "Synapse/internal/app/admin/rule/repository"
)

// service implementa a interface Service
type service struct {
	repo repository.Repository
}

// NewService cria uma nova instância do serviço de empresa
func NewService(r repository.Repository) Service {
	return &service{repo: r}
}

// GetAllRules retorna as regras paginadas (limit e offset)
func (s *service) FindAllRules(limit, offset int) ([]model.AdminRule, error) {
	return s.repo.FindAllRules(limit, offset)
}

// GetRuleByID retorna uma regra específica
func (s *service) FindRuleByID(ruleID int64) (*model.AdminRule, error) {
	return s.repo.FindRuleByID(ruleID)
}

// GetPermissionsByRuleID retorna os namespaces de permissões de uma regra
func (s *service) FindPermissionsByRuleID(ruleID int64) ([]string, error) {
	return s.repo.FindPermissionsByRuleID(ruleID)
}
