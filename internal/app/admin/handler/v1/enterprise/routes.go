package enterprise

import (
	controller "Synapse/internal/app/admin/enterprise/controller"
	repository "Synapse/internal/app/admin/enterprise/repository"
	service "Synapse/internal/app/admin/enterprise/service"
	db "Synapse/internal/database/db" // adaptável ao seu InitDB()

	"github.com/gin-gonic/gin"
)

func RegisterEnterpriseRoutes(router *gin.RouterGroup) {
	// Resolve dependências internas (DB > Repo > Service > Controller)
	dbConn := db.GetDB() // ou db.StartDB() — depende de sua implementação
	repo := repository.NewRepository(dbConn)
	svc := service.NewService(repo)
	ctrl := controller.NewEnterpriseController(svc)

	group := router.Group("/enterprise")
	{
		group.POST("/", ctrl.Create)
		group.GET("/", ctrl.ReadAll)
		group.GET("/cnpj/:cnpj", ctrl.ReadByCNPJ)
		group.GET("/nome/:nome", ctrl.ReadByNome)
		group.GET("/id/:id", ctrl.ReadByID)
		group.PUT("/cnpj/:cnpj", ctrl.UpdateByCNPJ)
	}
}
