package server

import (
	adminHandler "Synapse/internal/app/admin/handler"
	integrationHandler "Synapse/internal/app/integrations/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	adminHandler.RegisterAdminRoutes(router)
	integrationHandler.RegisterIntegrationsRoutes(router)
}
