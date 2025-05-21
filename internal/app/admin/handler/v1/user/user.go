package user

import (
	repositoryEnterprise "Synapse/internal/app/admin/enterprise/repository"
	adminMiddlewareRepository "Synapse/internal/app/admin/middleware/repository"
	adminMiddlewareService "Synapse/internal/app/admin/middleware/service"
	repositoryRule "Synapse/internal/app/admin/rule/repository"
	controller "Synapse/internal/app/admin/user/controller"
	repository "Synapse/internal/app/admin/user/repository"
	service "Synapse/internal/app/admin/user/service"
	rbac "Synapse/internal/app/middleware/auth/user"
	db "Synapse/internal/database/db"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	dbConn := db.GetDB()

	// Repositórios
	userRepo := repository.NewRepository(dbConn)
	enterpriseRepo := repositoryEnterprise.NewRepository(dbConn)
	ruleRepo := repositoryRule.NewRuleRepository(dbConn)
	adminRepo := adminMiddlewareRepository.NewMiddlewareRepository(dbConn)

	// Serviço de RBAC do admin
	adminMiddlewareService := adminMiddlewareService.NewMiddlewareService(adminRepo)

	// Middleware
	rbacMiddleware := rbac.NewRbacMiddleware(adminMiddlewareService)

	// Serviço
	svc := service.NewService(userRepo, enterpriseRepo, ruleRepo)

	// Controller
	ctrl := controller.NewUserController(svc)

	group := router.Group("/user")
	{
		group.POST("/", rbacMiddleware.RequirePermission("admin.user", "create"), ctrl.Create)
		group.POST("/token", ctrl.Token)
		group.GET("/:enterprise_id", rbacMiddleware.RequirePermission("admin.user", "read"), ctrl.ReadAll)
		group.GET("/email/:email", rbacMiddleware.RequirePermission("admin.user", "read"), ctrl.ReadByEmail)
		group.PUT("/:id", rbacMiddleware.RequirePermission("admin.user", "update"), ctrl.Update)
		group.DELETE("/:id", rbacMiddleware.RequirePermission("admin.user", "remove"), ctrl.Delete)
	}
}
