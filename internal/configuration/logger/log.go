package logger

import (
	"Synapse/internal/configuration/logger/log_print"
	"fmt"
)

func InicializaLog(env, level string) {
	fmt.Println("\nInicializando logger\n")
	fmt.Println("LOG ENV: ", env)
	fmt.Println("LOG LEVEL: ", level)
	// Log para print
	log_print.Init(env, level)
}
