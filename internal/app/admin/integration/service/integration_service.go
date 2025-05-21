package integration

import (
	enterpriseRepo "Synapse/internal/app/admin/enterprise/repository"
	model "Synapse/internal/app/admin/integration/model"
	repository "Synapse/internal/app/admin/integration/repository"
	"Synapse/internal/app/admin/pkg/security"
	userRepo "Synapse/internal/app/admin/user/repository"

	"context"
	"fmt"
)

type service struct {
	repo           repository.Repository
	enterpriseRepo enterpriseRepo.Repository
	userRepo       userRepo.Repository
}

func NewService(
	r repository.Repository,
	eRepo enterpriseRepo.Repository,
	uRepo userRepo.Repository,
) Service {
	return &service{
		repo:           r,
		enterpriseRepo: eRepo,
		userRepo:       uRepo,
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

// Remove a integração vinculada de uma empresa
func (s *service) DeleteIntegracaoFromEnterprise(enterpriseID, integracaoID int64) error {
	ctx := context.Background()

	// Verifica se a empresa existe
	if _, err := s.enterpriseRepo.ReadByID(enterpriseID); err != nil {
		return fmt.Errorf("empresa não encontrada")
	}

	// Verifica se a integração existe
	if _, err := s.repo.GetIntegracaoByID(ctx, integracaoID); err != nil {
		return fmt.Errorf("integração com ID %d não encontrada", integracaoID)
	}

	// Remove vínculo
	return s.repo.DeleteIntegracaoFromEnterprise(enterpriseID, integracaoID)
}

// Cria vinculo entre usuario e integracao
func (s *service) CreateIntegracaoUser(data model.IntegracaoUser) error {
	ctx := context.Background()

	// Verifica se a integração existe
	if _, err := s.repo.GetIntegracaoByID(ctx, data.IntegracaoID); err != nil {
		return fmt.Errorf("integração não encontrada")
	}

	// Verifica se o usuário existe
	if _, err := s.userRepo.ReadByID(data.UserID); err != nil {
		return fmt.Errorf("usuário não encontrado")
	}

	// Cria vínculo
	return s.repo.CreateIntegracaoUser(data)
}

// Criar token de integração
func (s *service) CreateTokenIntegracao(email, senha string, integracaoID int64) (string, error) {
	ctx := context.Background()

	user, err := s.userRepo.ValidateCredentials(ctx, email, senha)
	if err != nil {
		return "", fmt.Errorf("credenciais inválidas")
	}

	// Verifica se o usuário está vinculado à integração
	hasAccess, err := s.repo.CheckUserHasIntegracao(user.ID, integracaoID)
	if err != nil {
		return "", err
	}
	if !hasAccess {
		return "", fmt.Errorf("usuário não possui permissão para essa integração")
	}

	// Gera token simples (pode ser UUID, JWT ou outro hash)
	tokenData, err := security.GenerateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar token: %w", err)
	}

	// Salva no banco
	err = s.repo.SaveIntegracaoToken(user.ID, integracaoID, tokenData.Token)
	if err != nil {
		return "", err
	}

	return tokenData.Token, nil
}

// Busca  os dados de uma integração que o usuário possui
func (s *service) GetIntegracoesByUserID(userID int64) ([]model.IntegracaoUsuarioDetalhada, error) {
	// Valida se o usuário existe
	if _, err := s.userRepo.ReadByID(userID); err != nil {
		return nil, fmt.Errorf("usuário com ID %d não encontrado", userID)
	}

	return s.repo.GetIntegracoesByUserID(userID)
}

// RemoveIntegration remove uma integração do usuário, após validações
func (s *service) RemoveIntegrationFromUser(ctx context.Context, userID, integrationID int64) error {
	// Valida se o usuário existe
	user, err := s.userRepo.ReadByID(userID)
	if err != nil {
		return fmt.Errorf("erro ao buscar usuário: %w", err)
	}
	if user == nil {
		return fmt.Errorf("usuário %d não encontrado", userID)
	}

	// Valida se a integração existe
	integration, err := s.repo.GetIntegracaoByID(ctx, integrationID)
	if err != nil {
		return fmt.Errorf("erro ao buscar integração: %w", err)
	}
	if integration == nil {
		return fmt.Errorf("integração %d não encontrada", integrationID)
	}

	// Remove a integração do usuário
	if err := s.repo.RemoveIntegrationFromUser(ctx, userID, integrationID); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
