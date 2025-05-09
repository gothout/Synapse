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

	sort.Strings(migrationFiles)

	for _, file := range migrationFiles {
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
		dropSQL := convertCreateToDrop(string(sqlContent))
		fmt.Println("Comando utilizado: ", dropSQL)

		fmt.Println("🔄 Dropando:", file)
		if _, err := db.Exec(context.Background(), dropSQL); err != nil {
			return fmt.Errorf("erro ao executar drop %s: %w", file, err)
		}
		fmt.Println("🗑️ Tabela dropada com sucesso:", file)
	}

	return nil
}

// Função auxiliar para converter comandos CREATE TABLE em DROP TABLE
func convertCreateToDrop(sqlContent string) string {
	lines := strings.Split(sqlContent, "\n")
	var tableName string

	for _, line := range lines {
		upperLine := strings.ToUpper(strings.TrimSpace(line))

		if strings.HasPrefix(upperLine, "CREATE TABLE IF NOT EXISTS") {
			parts := strings.Fields(line)
			if len(parts) >= 6 {
				tableName = parts[5] // CREATE TABLE IF NOT EXISTS nome_tabela
			}
		} else if strings.HasPrefix(upperLine, "CREATE TABLE") {
			parts := strings.Fields(line)
			if len(parts) >= 3 {
				tableName = parts[2] // CREATE TABLE nome_tabela
			}
		}

		if tableName != "" {
			tableName = strings.Trim(tableName, `"';`)
			break
		}
	}

	return fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;", tableName)
}
