package cmd

import (
	cmdEnv "Synapse/cmd/configurations/env"
	cmdLog "Synapse/cmd/configurations/logger"
	"fmt"
)

func InitServer() {
	fmt.Println("Iniciando serviços\n")
	cmdEnv.InitEnv()
	cmdLog.InicializaLog()
	fmt.Println("Serviços iniciados")
}
