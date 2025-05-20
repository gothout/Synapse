package integration

import (
	controller "Synapse/internal/app/admin/integration/controller"
	repository "Synapse/internal/app/admin/integration/repository"
	service "Synapse/internal/app/admin/integration/service"
	adminMiddlewareRepository "Synapse/internal/app/admin/middleware/repository"
	adminMiddlewareService "Synapse/internal/app/admin/middleware/service"
	rbac "Synapse/internal/app/middleware/auth"
	db "Synapse/internal/database/db"

	"github.com/gin-gonic/gin"
)

func RegisterIntegrationRoutes(router *gin.RouterGroup) {
	// Conexão com o banco
	dbConn := db.GetDB()

	// Repositório e serviço da integração
	repo := repository.NewIntegrationRepository(dbConn)
	svc := service.NewService(repo)
	ctrl := controller.NewIntegrationController(svc)

	// Middleware de RBAC
	adminRepo := adminMiddlewareRepository.NewMiddlewareRepository(dbConn)
	adminService := adminMiddlewareService.NewMiddlewareService(adminRepo)
	rbacMiddleware := rbac.NewRbacMiddleware(adminService)

	// Grupo de rotas da integração
	group := router.Group("/integration")
	{
		group.GET("/", rbacMiddleware.RequirePermission("admin.integration", "read"), ctrl.GetAllIntegrations)
		group.GET("/marcas", rbacMiddleware.RequirePermission("admin.integration", "read"), ctrl.GetAllMarcas)
		group.GET("/marca/detalhada", rbacMiddleware.RequirePermission("admin.integration", "read"), ctrl.GetDetalhadasByMarcaID)
	}
}
