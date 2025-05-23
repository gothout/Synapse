package middleware

import (
	model "Synapse/internal/app/admin/middleware/model"
	repository "Synapse/internal/app/admin/middleware/repository"
	"context"
)

type middlewareService struct {
	repo repository.MiddlewareRepository
}

func NewMiddlewareService(repo repository.MiddlewareRepository) MiddlewareService {
	return &middlewareService{repo: repo}
}

// ValidateToken busca o ID do usuário a partir de um token, se válido
func (s *middlewareService) ValidateToken(ctx context.Context, token string) (int64, error) {
	return s.repo.GetUserIDByToken(ctx, token)
}

// HasPermission verifica se o usuário tem permissão para acessar determinado módulo e ação
func (s *middlewareService) HasPermission(ctx context.Context, userID int64, module, action string) (bool, error) {
	return s.repo.CheckPermission(ctx, userID, module, action)
}

// ValidateTokenIntegration retorna os dados da integração associada ao token
func (s *middlewareService) ValidateTokenIntegration(ctx context.Context, token string) (*model.IntegrationWithPermissions, error) {
	return s.repo.FindIntegrationByToken(ctx, token)
}

// CheckEnterpriseToken valida se o token fornecido está associado a um usuário e se esse usuário tem permissão
// para a integração solicitada, além de verificar se a empresa desse usuário também está associada à integração.
func (s *middlewareService) CheckEnterpriseToken(ctx context.Context, token string, integrationID string) (*model.IntegrationWithPermissions, error) {
	return s.repo.CheckEnterpriseToken(ctx, token, integrationID)
}
