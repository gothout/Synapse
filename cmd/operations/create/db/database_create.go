package db

import (
	dbPool "Synapse/internal/database/db"
	"Synapse/internal/database/migrations"
	"log"
)

type DatabaseCreator struct{}

func (DatabaseCreator) Run() {
	dbPool.StartDB()

	if err := migrations.RunAllMigrations(dbPool.GetDB()); err != nil {
		log.Fatalf("❌ Falha nas migrations: %v", err)
	}

	log.Println("✅ Migrations executadas com sucesso.")
}
