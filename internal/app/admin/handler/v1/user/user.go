package user

import (
	controller "Synapse/internal/app/admin/user/controller"
	repository "Synapse/internal/app/admin/user/repository"
	service "Synapse/internal/app/admin/user/service"

	repositoryEnterprise "Synapse/internal/app/admin/enterprise/repository"
	repositoryRule "Synapse/internal/app/admin/rule/repository"

	db "Synapse/internal/database/db"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	// Conexão com o banco
	dbConn := db.GetDB()

	// Repositórios
	userRepo := repository.NewRepository(dbConn)
	enterpriseRepo := repositoryEnterprise.NewRepository(dbConn)
	ruleRepo := repositoryRule.NewRuleRepository(dbConn)

	// Serviço
	svc := service.NewService(userRepo, enterpriseRepo, ruleRepo)

	// Controller
	ctrl := controller.NewUserController(svc)

	// Grupo de rotas
	group := router.Group("/user")
	{
		group.POST("/", ctrl.Create)
		group.GET("/:enterprise_id", ctrl.ReadAll)
		group.GET("/email/:email", ctrl.ReadByEmail)
		group.PUT("/:id", ctrl.Update)
	}
}
