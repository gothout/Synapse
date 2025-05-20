package server

import (
	"log"
	"os"

	adminHandler "Synapse/internal/app/admin/handler"
	corsMiddleware "Synapse/internal/app/middleware/server"

	"github.com/gin-gonic/gin"
)

// InitServer configura e retorna o router do Gin já com as rotas registradas.
func InitServer() *gin.Engine {
	router := gin.Default()
	Env := os.Getenv("ENV")
	if Env == "DEV" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// Configurando cors
	router.Use(corsMiddleware.MiddlewareCors())
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
