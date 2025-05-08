package db

import (
	db "Synapse/internal/database/db"
	"context"
	"fmt"
	"log"
)

// CheckDatabase é o tipo que implementa o check de conexão.
type CheckDatabase struct{}

func (CheckDatabase) Run() {
	db.StartDB()
	conn := db.GetDB()

	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("❌ Erro ao consultar versão do PostgreSQL: %v", err)
	}

	fmt.Printf("✅ PostgreSQL versão: %s\n", version)
}
