package cmd

import (
	cmdEnv "Synapse/cmd/configurations/env"
	cmdLog "Synapse/cmd/configurations/logger"
	cmdServer "Synapse/cmd/server"
	cmdDatabase "Synapse/cmd/validators"
	env "Synapse/internal/configuration/env"
	"fmt"
)

func InitServer() {
	fmt.Println("Iniciando serviços...\n")

	// 1. Inicializa env, log e validações
	cmdEnv.InitEnv()
	cmdLog.InicializaLog()
	cmdDatabase.ValidatorInicialize()

	fmt.Println("Serviços iniciados ✅")

	// 2. Inicializa o servidor
	router := cmdServer.InitServer()
	cmdServer.StartServer(router, env.GetPortServer())
}
