package middleware

import (
	"context"
	"fmt"
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

func (m *IntegrationRbacMiddleware) RequireIntegrationPermission(integrationID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))

		if token == "" {
			c.JSON(403, gin.H{"error": "Token de integração ausente"})
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
		// Valida o token e a permissão da empresa para a integração
		integration, err = m.Service.CheckEnterpriseToken(ctx, token, integration.ID)
		fmt.Print(integration)
		if err != nil || integration == nil {
			c.JSON(403, gin.H{"error": "Token inválido ou sem permissão para essa integração"})
			c.Abort()
			return
		}

		// ✅ Guarda os dados da integração no contexto
		c.Set("integration", integration)
		c.Next()
	}
}

///func normalize(s string) string {
//	return strings.ToLower(strings.TrimSpace(s))
//}
