package v1

import (
	enterprise "Synapse/internal/app/admin/handler/v1/enterprise"
	integration "Synapse/internal/app/admin/handler/v1/integration"
	rules "Synapse/internal/app/admin/handler/v1/rule"
	user "Synapse/internal/app/admin/handler/v1/user"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(router *gin.RouterGroup) {
	v1Group := router.Group("/v1")

	enterprise.RegisterEnterpriseRoutes(v1Group)
	rules.RegisterRuleRoutes(v1Group)
	user.RegisterUserRoutes(v1Group)
	integration.RegisterIntegrationRoutes(v1Group)
}
