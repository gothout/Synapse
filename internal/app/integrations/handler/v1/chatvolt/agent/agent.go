package agent

import (
	api "Synapse/internal/app/integrations/chatvolt/agent/api"
	controller "Synapse/internal/app/integrations/chatvolt/agent/controller"
	repository "Synapse/internal/app/integrations/chatvolt/agent/repository"
	service "Synapse/internal/app/integrations/chatvolt/agent/service"

	adminIntegrationRepo "Synapse/internal/app/admin/integration/repository"
	adminUserRepo "Synapse/internal/app/admin/user/repository"
	db "Synapse/internal/database/db"

	middlewareRepo "Synapse/internal/app/admin/middleware/repository"
	middlewareService "Synapse/internal/app/admin/middleware/service"
	rbac "Synapse/internal/app/middleware/auth/integration"

	"github.com/gin-gonic/gin"
)

// RegisterChatvoltRoutes configura as rotas do subdomínio /agent da Chatvolt
func RegisterChatvoltRoutes(router *gin.RouterGroup) {
	dbConn := db.GetDB()

	// Repositórios
	agentRepo := repository.NewAgentRepository(dbConn)
	integrationRepo := adminIntegrationRepo.NewIntegrationRepository(dbConn)
	userRepo := adminUserRepo.NewRepository(dbConn)
	middlewareRepo := middlewareRepo.NewMiddlewareRepository(dbConn)

	// API externa da Chatvolt
	chatvoltAPI := api.NewChatvoltAPI()

	// Serviços
	svc := service.NewAgentService(chatvoltAPI, agentRepo, integrationRepo, userRepo)
	middlewareSvc := middlewareService.NewMiddlewareService(middlewareRepo)

	// Middleware RBAC de integração
	rbacMiddleware := rbac.NewIntegrationRbacMiddleware(middlewareSvc)

	// Controller
	ctrl := controller.NewAgentController(svc)

	// Grupo: /agent
	group := router.Group("/agent")
	{
		group.POST("/config", rbacMiddleware.RequireIntegrationPermission("agent"), ctrl.PostAgentConfig)
	}
}
