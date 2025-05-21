package middleware

import (
	"context"
	"strings"

	adminService "Synapse/internal/app/admin/middleware/service"
	"Synapse/internal/configuration/rest_err"

	"github.com/gin-gonic/gin"
)

type IntegrationRbacMiddleware struct {
	Service adminService.MiddlewareService
}

func NewIntegrationRbacMiddleware(service adminService.MiddlewareService) *IntegrationRbacMiddleware {
	return &IntegrationRbacMiddleware{Service: service}
}

func (m *IntegrationRbacMiddleware) RequireIntegrationPermission(requiredScope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))

		if token == "" {
			err := rest_err.NewForbiddenError("Token de integração ausente")
			c.JSON(err.Code, err)
			c.Abort()
			return
		}

		ctx := context.Background()
		integration, err := m.Service.ValidateTokenIntegration(ctx, token)
		if err != nil || integration == nil {
			res := rest_err.NewForbiddenError("Token de integração inválido ou não autorizado")
			c.JSON(res.Code, res)
			c.Abort()
			return
		}

		// Comparação baseada no nome da integração
		if normalize(integration.Nome) != normalize(requiredScope) {
			res := rest_err.NewForbiddenError("Token de integração inválido ou não autorizado")
			//res := rest_err.NewForbiddenError("Integração sem permissão para: " + requiredScope)
			c.JSON(res.Code, res)
			c.Abort()
			return
		}

		// ✅ Permissão concedida com base no nome da integração
		c.Set("integration", integration)
		c.Next()
	}
}

func normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
