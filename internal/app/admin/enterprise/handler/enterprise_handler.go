package enterprise

import (
	controller "Synapse/internal/app/admin/enterprise/controller"
	service "Synapse/internal/app/admin/enterprise/service"

	"github.com/gin-gonic/gin"
)

func RegisterEnterpriseRoutes(router *gin.RouterGroup, service service.Service) {
	enterpriseController := controller.NewEnterpriseController(service)

	enterpriseGroup := router.Group("/enterprise")
	{
		enterpriseGroup.POST("/", enterpriseController.Create)
	}
}
