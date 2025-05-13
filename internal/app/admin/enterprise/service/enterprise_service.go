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

// Busca empresa por CNPJ
func (s *service) ReadByCNPJ(cnpj string) (*model.AdminEnterprise, error) {
	return s.repo.ReadByCNPJ(cnpj)
}

// Busca empresa por NOME
func (s *service) ReadByNome(nome string) (*model.AdminEnterprise, error) {
	return s.repo.ReadByNome(nome)
}

// Buscar empresa por ID
func (s *service) ReadByID(id int64) (*model.AdminEnterprise, error) {
	return s.repo.ReadByID(id)
}

// Buscar empresas por Page
func (s *service) ReadAll(page int) ([]model.AdminEnterprise, error) {
	return s.repo.ReadAll(page)
}

// Atualiza empresa por CNPJ
func (s *service) UpdateByCNPJ(cnpj, newCNPJ string, enterprise *model.AdminEnterprise) (*model.AdminEnterprise, error) {
	return s.repo.UpdateByCNPJ(cnpj, newCNPJ, enterprise)
}
