package v1

import (
	chatvolt "Synapse/internal/app/integrations/handler/v1/chatvolt"

	"github.com/gin-gonic/gin"
)

// RegisterV1Routes cria o grupo /v1 e carrega os domínios filhos
func RegisterV1Routes(router *gin.RouterGroup) {
	v1Group := router.Group("/v1")
	chatvolt.RegisterV1Routes(v1Group)
}
