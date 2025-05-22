package chatvolt

import (
	agentChatVolt "Synapse/internal/app/integrations/handler/v1/chatvolt/agent"

	"github.com/gin-gonic/gin"
)

// RegisterV1Routes cria o grupo /chatvolt e registra subdomínio agent
func RegisterV1Routes(router *gin.RouterGroup) {
	chatvoltGroup := router.Group("/chatvolt")
	agentChatVolt.RegisterChatvoltRoutes(chatvoltGroup)
}
