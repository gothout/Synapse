package check

import (
	dbPool "Synapse/internal/database/db"
	"Synapse/internal/database/migrations"
	"log"
	"os"
)

type DatabaseDropper struct{}

func (DatabaseDropper) Run() {
	dbPool.StartDB()

	if err := migrations.DropAllTables(os.Getenv("DATABASE_PASSWORD"), dbPool.GetDB()); err != nil {
		log.Fatalf("❌ Falha no drop das tabelas: %v", err)
	}

}
