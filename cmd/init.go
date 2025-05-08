package cmd

import (
	cmdEnv "Synapse/cmd/configurations/env"
	cmdLog "Synapse/cmd/configurations/logger"
	cmdDatabase "Synapse/cmd/validators"
	"fmt"
)

func InitServer() {
	fmt.Println("Iniciando serviços\n")
	cmdEnv.InitEnv()
	cmdLog.InicializaLog()
	// Efetuando validações de banco.
	cmdDatabase.ValidatorInicialize()
	fmt.Println("Serviços iniciados")
}
