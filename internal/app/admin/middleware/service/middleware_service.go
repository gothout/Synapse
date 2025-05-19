package middleware

import (
	repository "Synapse/internal/app/admin/middleware/repository"
	"context"
)

type middlewareService struct {
	repo repository.MiddlewareRepository
}

func NewMiddlewareService(repo repository.MiddlewareRepository) MiddlewareService {
	return &middlewareService{repo: repo}
}

func (s *middlewareService) ValidateToken(ctx context.Context, token string) (int64, error) {
	return s.repo.GetUserIDByToken(ctx, token)
}

func (s *middlewareService) HasPermission(ctx context.Context, userID int64, module, action string) (bool, error) {
	return s.repo.CheckPermission(ctx, userID, module, action)
}
