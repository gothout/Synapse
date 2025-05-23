package handler

import (
	_ "Synapse/docs" // esse import embute os comentários do swag
	v1 "Synapse/internal/app/admin/handler/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterAdminRoutes registra todas as rotas da área administrativa
func RegisterAdminRoutes(router *gin.Engine) {
	adminGroup := router.Group("/admin")
	// Inicia Swagger.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1.RegisterV1Routes(adminGroup)
}
