package log_print

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	isDev    bool
	logLevel string
)

// Init configura o logger do sistema com base nas variáveis de ambiente:
// ENV  → "DEV", "PROD", etc.
// LOG  → "DEBUG", "INFO", "WARNING", "ERROR"
func Init(env string, level string) {
	isDev = (env == "DEV" || env == "local")
	logLevel = strings.ToUpper(level)

	if !isDev {
		date := time.Now().Format("2006-01-02")
		logDir := "logs"
		_ = os.MkdirAll(logDir, os.ModePerm)

		filename := fmt.Sprintf("%s/app_%s_%s.log", logDir, env, date)

		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("❌ Não foi possível criar o arquivo de log: %v", err)
		}

		log.SetOutput(file)
		log.SetFlags(log.LstdFlags | log.Llongfile) // Caminho completo do arquivo
	}
}

// Debug exibe logs apenas se LOG=DEBUG
func Debug(message string) {
	if logLevel == "DEBUG" {
		log.Println("[DEBUG]", message)
	}
}

// Info exibe logs se LOG=DEBUG ou INFO
func Info(message string) {
	if logLevel == "DEBUG" || logLevel == "INFO" {
		log.Println("[INFO]", message)
	}
}

// Warn exibe logs se LOG=DEBUG, INFO ou WARNING
func Warn(message string) {
	if logLevel == "DEBUG" || logLevel == "INFO" || logLevel == "WARNING" {
		log.Println("[WARN]", message)
	}
}

// Error exibe logs se LOG=DEBUG, INFO, WARNING ou ERROR
func Error(err error) {
	if logLevel == "DEBUG" || logLevel == "INFO" || logLevel == "WARNING" || logLevel == "ERROR" {
		log.Println("[ERROR]", err)
	}
}

// Fatal sempre encerra o sistema com log
func Fatal(err error) {
	log.Fatal("[FATAL]", err)
}
