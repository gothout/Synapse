package env

import (
	"fmt"
	"log"
	"os"
)

func Configuration() {
	fmt.Println("\nIniciando variaveis de ambiente.")
	LoadEnv()
	// Verificando ENV
	Env := os.Getenv("ENV")
	fmt.Println("Env: ", Env)

	switch Env {
	case "DEV":
		fmt.Println("Voce está rodando Env de desenvolvimento!")
		InitEnvsDev()
	case "PROD":
		fmt.Println("Voce está rodando Env de produção!")
		InitEnvProd()
	default:
		log.Fatal("Defina qual env você está rodando")
		os.Exit(1)
	}

}

func InitEnvsDev() {

	fmt.Println("Iniciando Envs de Servidor")
	fmt.Println("Host do Servidor: ", GetHostServer())
	fmt.Println("Porta do Servidor: ", GetPortServer())

	fmt.Println("Iniciando Envs de Banco de Dados")
	fmt.Println("Host do Banco de Dados: ", GetDatabaseHost())
	fmt.Println("Porta do Banco de Dados: ", GetDatabasePort())
	fmt.Println("Nome do Banco de Dados: ", GetDatabaseName())
	fmt.Println("Usuário do Banco de Dados: ", GetDatabaseUser())
	fmt.Println("Senha do Banco de Dados: ", GetDatabasePassword())

	fmt.Println("Iniciando Envs de log")
	fmt.Println("Env de log sistematica: ", GetLog())
	// Buscando secretkey sem logar.
	GetSecretKey()
}

func InitEnvProd() {

	fmt.Println("Iniciando Envs de Servidor")
	fmt.Println("Host do Servidor: ", GetHostServer())
	fmt.Println("Porta do Servidor: ", GetPortServer())

	fmt.Println("Iniciando Envs de Banco de Dados")
	fmt.Println("Host do Banco de Dados: ", GetDatabaseHost())
	fmt.Println("Porta do Banco de Dados: ", GetDatabasePort())
	fmt.Println("Nome do Banco de Dados: ", GetDatabaseName())
	fmt.Println("Usuário do Banco de Dados: ", GetDatabaseUser())
	fmt.Println("Senha do Banco de Dados: ", GetDatabasePassword())

	fmt.Println("Iniciando Envs de log")
	fmt.Println("Env de log sistematica: ", GetLog())
	// Buscando secretkey sem logar.
	GetSecretKey()

}
