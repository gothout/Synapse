package env

import (
	"log"
	"os"
)

func GetLog() string {
	if os.Getenv("LOG") == "" || os.Getenv("LOG") != "DEBUG" && os.Getenv("LOG") != "INFO" && os.Getenv("LOG") != "ERROR" && os.Getenv("LOG") != "" {
		log.Fatal("Env LOG nao setado ou com valor invalido")
		os.Exit(1)
	}
	return os.Getenv("LOG")
}
