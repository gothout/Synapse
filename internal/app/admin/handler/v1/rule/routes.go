package rule

import (
	controller "Synapse/internal/app/admin/rule/controller"
	repository "Synapse/internal/app/admin/rule/repository"
	service "Synapse/internal/app/admin/rule/service"
	db "Synapse/internal/database/db"

	"github.com/gin-gonic/gin"
)

func RegisterRuleRoutes(router *gin.RouterGroup) {
	// Resolve dependências internas (DB > Repo > Service > Controller)
	dbConn := db.GetDB()
	repo := repository.NewRuleRepository(dbConn)
	svc := service.NewService(repo)
	ctrl := controller.NewRuleController(svc)

	group := router.Group("/rules")
	{
		group.GET("/", ctrl.GetAll)
		group.GET("/:id", ctrl.GetByID)
		group.GET("/:id/permissions", ctrl.GetPermissions)
	}
}
