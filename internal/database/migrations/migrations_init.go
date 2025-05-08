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
