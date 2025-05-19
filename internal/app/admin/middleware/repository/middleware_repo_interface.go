package middleware

import (
	"context"
)

type MiddlewareRepository interface {
	GetUserIDByToken(ctx context.Context, token string) (int64, error)
	CheckPermission(ctx context.Context, userID int64, module string, action string) (bool, error)
}
