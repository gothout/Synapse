package integration

import (
	enterpriseRepo "Synapse/internal/app/admin/enterprise/repository"
	controller "Synapse/internal/app/admin/integration/controller"
	repository "Synapse/internal/app/admin/integration/repository"
	service "Synapse/internal/app/admin/integration/service"
	adminMiddlewareRepository "Synapse/internal/app/admin/middleware/repository"
	adminMiddlewareService "Synapse/internal/app/admin/middleware/service"
	userRepo "Synapse/internal/app/admin/user/repository"
	rbac "Synapse/internal/app/middleware/auth"
	db "Synapse/internal/database/db"

	"github.com/gin-gonic/gin"
)

func RegisterIntegrationRoutes(router *gin.RouterGroup) {
	dbConn := db.GetDB()

	// Repositórios
	repo := repository.NewIntegrationRepository(dbConn)
	entRepo := enterpriseRepo.NewRepository(dbConn)
	userRepo := userRepo.NewRepository(dbConn)

	// Serviço
	svc := service.NewService(repo, entRepo, userRepo)
	ctrl := controller.NewIntegrationController(svc)

	// Middleware
	adminRepo := adminMiddlewareRepository.NewMiddlewareRepository(dbConn)
	adminService := adminMiddlewareService.NewMiddlewareService(adminRepo)
	rbacMiddleware := rbac.NewRbacMiddleware(adminService)

	// Rotas
	group := router.Group("/integration")
	{
		group.GET("/", rbacMiddleware.RequirePermission("admin.integration", "read"), ctrl.GetAllIntegrations)
		group.GET("/marcas", rbacMiddleware.RequirePermission("admin.integration", "read"), ctrl.GetAllMarcas)
		group.GET("/marca/detalhada", rbacMiddleware.RequirePermission("admin.integration", "read"), ctrl.GetDetalhadasByMarcaID)
		group.POST("/enterprise", rbacMiddleware.RequirePermission("admin.integration", "create"), ctrl.CreateIntegracaoEnterprise)
		group.GET("/enterprise/:enterprise_id", rbacMiddleware.RequirePermission("admin.integration", "read"), ctrl.GetByEnterpriseID)
		group.DELETE("/enterprise", rbacMiddleware.RequirePermission("admin.integration", "remove"), ctrl.DeleteIntegracaoFromEnterprise)
		group.POST("/user", rbacMiddleware.RequirePermission("admin.user", "create"), ctrl.CreateIntegracaoUser)
		group.POST("/token", rbacMiddleware.RequirePermission("admin.integration", "create"), ctrl.CreateTokenIntegracao)
	}
}
