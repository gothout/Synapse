package database

import (
	Print "Synapse/internal/configuration/logger/log_print"
	database "Synapse/internal/database/db"
	"context"
	"fmt"
)

// Init inicia a conex o com o banco de dados PostgreSQL.
//
// Esta fun o   usada para verificar se o banco de dados est  dispon vel e
// quaisquer problemas de conex o s o tratados aqui.
//
// Retorna um erro caso n o consiga conectar ao banco de dados.
func Init() error {
	database.StartDB()
	db := database.GetDB()

	var version string
	if err := db.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		return fmt.Errorf("erro ao consultar versão do PostgreSQL: %w", err)
	}
	Print.Info(fmt.Sprintf("✅ PostgreSQL versão: %s", version))
	return nil
}
