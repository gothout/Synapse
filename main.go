// @title           Synapse API
// @version         1.0
// @description     API administrativa do sistema Synapse

// 🔐 Definição do Bearer Token
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Forneça o token no formato: Bearer <token>

package main

import (
	"fmt"
	"os"
	"strings"

	"Synapse/cmd"
	checks "Synapse/cmd/operations"
)

func main() {

	// Se o primeiro argumento for um comando tipo --check-db, executa o handler
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "--") {
		arg := strings.ToLower(os.Args[1])
		if checker, ok := checks.Get(arg); ok {
			checker.Run()
			return
		}

		fmt.Printf("❌ Comando não reconhecido: %s\n digite --help para ajuda.", arg)
		return
	}

	// Se não recebeu argumento especial, sobe o servidor
	cmd.InitServer()
}
