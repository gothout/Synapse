package middleware

import (
	model "Synapse/internal/app/admin/middleware/model"
	"context"
)

type MiddlewareService interface {
	ValidateToken(ctx context.Context, token string) (int64, error)
	HasPermission(ctx context.Context, userID int64, module, action string) (bool, error)
	ValidateTokenIntegration(ctx context.Context, token string) (*model.IntegrationWithPermissions, error)
	CheckEnterpriseToken(ctx context.Context, token string, integrationID string) (*model.IntegrationWithPermissions, error)
}
