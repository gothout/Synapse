package middleware

import (
	"context"
)

type MiddlewareService interface {
	ValidateToken(ctx context.Context, token string) (int64, error)
	HasPermission(ctx context.Context, userID int64, module, action string) (bool, error)
}
