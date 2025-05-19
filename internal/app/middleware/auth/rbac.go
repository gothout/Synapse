package middleware

import (
	"context"
	"strings"

	adminService "Synapse/internal/app/admin/middleware/service"
	"Synapse/internal/configuration/rest_err"

	"github.com/gin-gonic/gin"
)

type RbacMiddleware struct {
	Service adminService.MiddlewareService
}

func NewRbacMiddleware(service adminService.MiddlewareService) *RbacMiddleware {
	return &RbacMiddleware{Service: service}
}

func (m *RbacMiddleware) RequirePermission(module, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")

		if token == "" {
			err := rest_err.NewForbiddenError("Token ausente")
			c.JSON(err.Code, err)
			c.Abort()
			return
		}

		ctx := context.Background()
		userID, err := m.Service.ValidateToken(ctx, token)
		if err != nil {
			res := rest_err.NewForbiddenError("Token inválido ou expirado")
			c.JSON(res.Code, res)
			c.Abort()
			return
		}

		allowed, err := m.Service.HasPermission(ctx, userID, module, action)
		if err != nil {
			res := rest_err.NewInternalServerError("Erro ao verificar permissão", []rest_err.Causes{{Message: err.Error()}})
			c.JSON(res.Code, res)
			c.Abort()
			return
		}

		if !allowed {
			res := rest_err.NewForbiddenError("Permissão negada para " + module + "." + action)
			c.JSON(res.Code, res)
			c.Abort()
			return
		}

		// Permissão concedida
		c.Set("user_id", userID)
		c.Next()
	}
}
