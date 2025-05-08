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

func Init(env string, level string) {
	isDev = (env == "DEV" || env == "local")
	logLevel = strings.ToUpper(level)

	// Se NÃO estiver em DEV, redireciona log para arquivo
	if !isDev {
		date := time.Now().Format("2006-01-02")
		filename := fmt.Sprintf("log_%s_%s.log", env, date)

		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("❌ Não foi possível criar o arquivo de log: %v", err)
		}

		log.SetOutput(file)
	}
}

func Debug(message string) {
	if logLevel == "DEBUG" {
		log.Println("[DEBUG]", message)
	}
}

func Info(message string) {
	if logLevel == "DEBUG" || logLevel == "INFO" {
		log.Println("[INFO]", message)
	}
}

func Warn(message string) {
	if logLevel == "DEBUG" || logLevel == "INFO" || logLevel == "WARNING" {
		log.Println("[WARN]", message)
	}
}

func Error(err error) {
	if logLevel == "DEBUG" || logLevel == "INFO" || logLevel == "WARNING" || logLevel == "ERROR" {
		log.Println("[ERROR]", err)
	}
}

func Fatal(err error) {
	log.Fatal("[FATAL]", err)
}

/* EXEMPLO DE USO
import (
	"os"
	"synapse/internal/configuration/logger/log_print"
)

func main() {
	env := os.Getenv("ENV")
	level := os.Getenv("LOG")

	log_print.Init(env, level)

	log_print.Debug("Iniciando sistema")
	log_print.Info("Sistema rodando")
	log_print.Warn("Essa é um alerta")
	log_print.Error(fmt.Errorf("algo deu errado"))
}
*/
