package integration

import (
	enterpriseRepo "Synapse/internal/app/admin/enterprise/repository"
	model "Synapse/internal/app/admin/integration/model"
	repository "Synapse/internal/app/admin/integration/repository"
	"context"
	"fmt"
)

type service struct {
	repo           repository.Repository
	enterpriseRepo enterpriseRepo.Repository
}

// NewService cria uma nova instância do serviço de integração
func NewService(
	r repository.Repository,
	eRepo enterpriseRepo.Repository,
) Service {
	return &service{
		repo:           r,
		enterpriseRepo: eRepo,
	}
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

// Cria a relação entre empresa e integração
func (s *service) CreateIntegracaoEnterprise(data model.IntegracaoEnterprise) error {
	ctx := context.Background()

	// Verifica se a empresa existe
	if _, err := s.enterpriseRepo.ReadByID(data.EnterpriseId); err != nil {
		return fmt.Errorf("empresa com ID não encontrada")
	}

	// Verifica se a integração existe
	if _, err := s.repo.GetIntegracaoByID(ctx, data.IntegracaoId); err != nil {
		return fmt.Errorf("integração com ID %d não encontrada", data.IntegracaoId)
	}

	// Vincula a integração à empresa
	return s.repo.CreateIntegracaoEnterprise(data)
}

// Retorna detalhe de integrações para X empresa
func (s *service) GetIntegracoesByEnterpriseID(enterpriseID int64) ([]model.IntegracaoEmpresaDetalhada, error) {
	// Verifica se a empresa existe
	_, err := s.enterpriseRepo.ReadByID(enterpriseID)
	if err != nil {
		return nil, fmt.Errorf("empresa não encontrada")
	}

	// Busca integrações liberadas para a empresa
	return s.repo.GetIntegracoesByEnterpriseID(enterpriseID)
}
