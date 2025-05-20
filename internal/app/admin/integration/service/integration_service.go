package integration

import (
	model "Synapse/internal/app/admin/integration/model"
	repository "Synapse/internal/app/admin/integration/repository"
)

// service implementa a interface Service
type service struct {
	repo repository.Repository
}

// NewService cria uma nova instância do serviço de empresa
func NewService(r repository.Repository) Service {
	return &service{repo: r}
}

// Retorna todas as integrações
func (s *service) GetAllIntegrations() ([]model.Integration, error) {
	return s.repo.GetAllIntegrations()
}

// Retorna todas as marcas donas das integrações
func (s *service) GetAllMarcas() ([]model.Marca, error) {
	return s.repo.GetAllMarcas()
}

// Buscar todas as integrações baseado no MarcaID
func (s *service) GetIntegrationsByMarcaID(marcaID int64) ([]model.IntegracaoDetalhada, error) {
	return s.repo.GetIntegracoesDetalhadasByMarcaID(marcaID)
}
