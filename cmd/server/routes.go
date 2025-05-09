package server

import (
	adminHandler "Synapse/internal/app/admin/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	adminHandler.RegisterAdminRoutes(router)
}
