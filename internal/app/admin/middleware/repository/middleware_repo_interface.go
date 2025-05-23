package middleware

import (
	model "Synapse/internal/app/admin/middleware/model"
	"context"
)

type MiddlewareRepository interface {
	GetUserIDByToken(ctx context.Context, token string) (int64, error)
	CheckPermission(ctx context.Context, userID int64, module, action string) (bool, error)
	FindIntegrationByToken(ctx context.Context, token string) (*model.IntegrationWithPermissions, error)
	CheckEnterpriseToken(ctx context.Context, token string, nomeIntegracao string) (*model.IntegrationWithPermissions, error)
}
