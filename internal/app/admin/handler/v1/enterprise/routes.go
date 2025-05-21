package enterprise

import (
	controller "Synapse/internal/app/admin/enterprise/controller"
	repository "Synapse/internal/app/admin/enterprise/repository"
	service "Synapse/internal/app/admin/enterprise/service"
	adminMiddlewareRepository "Synapse/internal/app/admin/middleware/repository"
	adminMiddlewareService "Synapse/internal/app/admin/middleware/service"
	rbac "Synapse/internal/app/middleware/auth/user"
	db "Synapse/internal/database/db" // adaptável ao seu InitDB()

	"github.com/gin-gonic/gin"
)

func RegisterEnterpriseRoutes(router *gin.RouterGroup) {
	// Resolve dependências internas (DB > Repo > Service > Controller)
	dbConn := db.GetDB() // ou db.StartDB() — depende de sua implementação
	repo := repository.NewRepository(dbConn)
	svc := service.NewService(repo)
	ctrl := controller.NewEnterpriseController(svc)
	adminRepo := adminMiddlewareRepository.NewMiddlewareRepository(dbConn)

	// Serviço de RBAC do admin
	adminMiddlewareService := adminMiddlewareService.NewMiddlewareService(adminRepo)

	// Middleware
	rbacMiddleware := rbac.NewRbacMiddleware(adminMiddlewareService)

	group := router.Group("/enterprise")
	{
		group.POST("/", rbacMiddleware.RequirePermission("admin.enterprise", "create"), ctrl.Create)
		group.GET("/", rbacMiddleware.RequirePermission("admin.enterprise", "read"), ctrl.ReadAll)
		group.GET("/cnpj/:cnpj", rbacMiddleware.RequirePermission("admin.enterprise", "read"), ctrl.ReadByCNPJ)
		group.GET("/nome/:nome", rbacMiddleware.RequirePermission("admin.enterprise", "read"), ctrl.ReadByNome)
		group.GET("/id/:id", rbacMiddleware.RequirePermission("admin.enterprise", "read"), ctrl.ReadByID)
		group.PUT("/cnpj/:cnpj", rbacMiddleware.RequirePermission("admin.enterprise", "update"), ctrl.UpdateByCNPJ)
		group.DELETE("/cnpj/:cnpj", rbacMiddleware.RequirePermission("admin.enterprise", "delete"), ctrl.DeleteByCNPJ)
	}
}
