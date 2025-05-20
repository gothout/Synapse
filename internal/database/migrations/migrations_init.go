package migrations

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RunAllMigrations(db *pgxpool.Pool) error {
	baseDir := "internal/database/migrations"

	// Ordem definida das pastas de migração
	migrationOrder := []string{
		"01_enterprise",
		"02_rules",
		"03_user",
		"04_integration",
	}

	for _, folder := range migrationOrder {
		fullPath := filepath.Join(baseDir, folder)

		files, err := os.ReadDir(fullPath)
		if err != nil {
			return fmt.Errorf("erro ao ler pasta %s: %w", folder, err)
		}

		// Ordena os arquivos da pasta
		var sqlFiles []string
		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
				sqlFiles = append(sqlFiles, filepath.Join(fullPath, file.Name()))
			}
		}
		sort.Strings(sqlFiles)

		// Executa os arquivos em ordem
		for _, file := range sqlFiles {
			sqlContent, err := os.ReadFile(file)
			if err != nil {
				return fmt.Errorf("erro ao ler %s: %w", file, err)
			}

			fmt.Println("🔄 Executando:", file)
			if _, err := db.Exec(context.Background(), string(sqlContent)); err != nil {
				return fmt.Errorf("erro ao executar %s: %w", file, err)
			}
			fmt.Println("✅ Sucesso:", file)
		}
	}

	return nil
}

func DropAllTables(password string, db *pgxpool.Pool) error {

	if password != os.Getenv("DATABASE_PASSWORD") {
		return fmt.Errorf("senha incorreta")
	}

	baseDir := "internal/database/migrations"
	var migrationFiles []string
	// Busca todos arquivos .sql recursivamente
	err := filepath.Walk(baseDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".sql") {
			migrationFiles = append(migrationFiles, path)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("erro ao buscar migrations: %w", err)
	}

	// Inverte a ordem dos arquivos para dropar corretamente.
	sort.Sort(sort.Reverse(sort.StringSlice(migrationFiles)))

	for _, file := range migrationFiles {
		sqlContent, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("erro ao ler %s: %w", file, err)
		}

		// Modifica o comando SQL CREATE TABLE para DROP TABLE
		dropSQLs := convertCreateToDrop(string(sqlContent))
		for _, dropSQL := range dropSQLs {
			fmt.Println("Comando utilizado: ", dropSQL)
			if _, err := db.Exec(context.Background(), dropSQL); err != nil {
				fmt.Println("⚠️ Erro ao executar drop:", err)
				continue // continua mesmo com erro
			}
			fmt.Println("🗑️ Tabela dropada com sucesso.")
		}

	}

	return nil
}

// Função auxiliar para converter comandos CREATE TABLE em DROP TABLE
func convertCreateToDrop(sqlContent string) []string {
	lines := strings.Split(sqlContent, "\n")
	var dropCommands []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		upper := strings.ToUpper(trimmed)

		if strings.HasPrefix(trimmed, "--") || trimmed == "" {
			continue
		}

		if strings.HasPrefix(upper, "CREATE TABLE") {
			// Remove IF NOT EXISTS se existir
			trimmed = strings.ReplaceAll(trimmed, "IF NOT EXISTS", "")
			trimmed = strings.ReplaceAll(trimmed, "CREATE TABLE", "")
			trimmed = strings.TrimSpace(trimmed)

			// Extrai nome da tabela
			tableName := strings.Fields(trimmed)[0]
			tableName = strings.Trim(tableName, `"'();`)

			drop := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;", tableName)
			dropCommands = append(dropCommands, drop)
		}
	}

	return dropCommands
}
