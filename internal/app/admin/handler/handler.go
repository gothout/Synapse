package handler

import (
	v1 "Synapse/internal/app/admin/handler/v1"

	"github.com/gin-gonic/gin"
)

// RegisterAdminRoutes registra todas as rotas da área administrativa
func RegisterAdminRoutes(router *gin.Engine) {
	adminGroup := router.Group("/admin")

	v1.RegisterV1Routes(adminGroup)
}
