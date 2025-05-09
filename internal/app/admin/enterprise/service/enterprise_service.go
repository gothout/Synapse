package enterprise

import (
	model "Synapse/internal/app/admin/enterprise/model"
	repository "Synapse/internal/app/admin/enterprise/repository"
)

// service implementa a interface Service
type service struct {
	repo repository.Repository
}

// NewService cria uma nova instância do serviço de empresa
func NewService(r repository.Repository) Service {
	return &service{repo: r}
}

// Create chama o repositório para persistir a empresa
func (s *service) Create(enterprise *model.AdminEnterprise) (*model.AdminEnterprise, error) {
	return s.repo.Create(enterprise)
}
