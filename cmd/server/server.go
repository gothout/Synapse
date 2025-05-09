package server

import (
	"log"

	adminHandler "Synapse/internal/app/admin/handler"

	"github.com/gin-gonic/gin"
)

// InitServer configura e retorna o router do Gin já com as rotas registradas.
func InitServer() *gin.Engine {
	router := gin.Default()

	// Aqui registramos todos os domínios, cada um resolve suas dependências
	adminHandler.RegisterAdminRoutes(router)

	// futuramente: publicHandler.RegisterPublicRoutes(router)
	return router
}

// StartServer inicia o servidor HTTP com o router e porta informados.
func StartServer(router *gin.Engine, port string) {
	log.Printf("🚀 Servidor iniciado em http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Falha ao iniciar servidor: %v", err)
	}
}
