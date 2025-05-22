package handler

import (
	_ "Synapse/docs"
	v1 "Synapse/internal/app/integrations/handler/v1"

	"github.com/gin-gonic/gin"
)

// RegisterIntegrationsRoutes registra o grupo /integrations
func RegisterIntegrationsRoutes(router *gin.Engine) {
	integrationGroup := router.Group("/integrations")
	v1.RegisterV1Routes(integrationGroup)
}
