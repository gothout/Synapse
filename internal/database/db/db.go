package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"Synapse/internal/configuration/env"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func StartDB() {
	env.LoadEnv()

	user := env.GetDatabaseUser()
	password := url.QueryEscape(env.GetDatabasePassword()) // Encode automático
	host := env.GetDatabaseHost()
	port := env.GetDatabasePort()
	database := env.GetDatabaseName()

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	db, err = pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("❌ Erro ao criar pool de conexão: %v", err)
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatalf("❌ Banco de dados indisponível: %v", err)
	}
}

func GetDB() *pgxpool.Pool {
	if db == nil {
		log.Fatal("⚠️ Banco de dados ainda não iniciado. Chame StartDB() primeiro.")
	}
	return db
}
