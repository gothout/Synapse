package rule

import (
	adminMiddlewareRepository "Synapse/internal/app/admin/middleware/repository"
	adminMiddlewareService "Synapse/internal/app/admin/middleware/service"
	controller "Synapse/internal/app/admin/rule/controller"
	repository "Synapse/internal/app/admin/rule/repository"
	service "Synapse/internal/app/admin/rule/service"
	rbac "Synapse/internal/app/middleware/auth"
	db "Synapse/internal/database/db"

	"github.com/gin-gonic/gin"
)

func RegisterRuleRoutes(router *gin.RouterGroup) {
	// Resolve dependências internas (DB > Repo > Service > Controller)
	dbConn := db.GetDB()
	repo := repository.NewRuleRepository(dbConn)
	svc := service.NewService(repo)
	ctrl := controller.NewRuleController(svc)
	adminRepo := adminMiddlewareRepository.NewMiddlewareRepository(dbConn)

	// Serviço de RBAC do admin
	adminMiddlewareService := adminMiddlewareService.NewMiddlewareService(adminRepo)

	// Middleware
	rbacMiddleware := rbac.NewRbacMiddleware(adminMiddlewareService)
	group := router.Group("/rules")
	{
		group.GET("/", rbacMiddleware.RequirePermission("admin.rules", "read"), ctrl.GetAll)
		group.GET("/:id", rbacMiddleware.RequirePermission("admin.rules", "read"), ctrl.GetByID)
		group.GET("/:id/permissions", rbacMiddleware.RequirePermission("admin.rules", "read"), ctrl.GetPermissions)
	}
}
