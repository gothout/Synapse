package logger

import (
	"Synapse/internal/configuration/env"
	"Synapse/internal/configuration/logger"
	"os"
)

func InicializaLog() {
	logger.InicializaLog(os.Getenv("ENV"), env.GetLog())
}
