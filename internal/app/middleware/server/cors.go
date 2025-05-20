package server

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// MiddlewareCors retorna o middleware de CORS configurado
func MiddlewareCors() gin.HandlerFunc {
	origins := []string{"http://localhost:8081"}
	if os.Getenv("ENV") == "PROD" {
		origins = []string{"https://ws.wonit.net.br"}
	}

	config := cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	return cors.New(config)
}
