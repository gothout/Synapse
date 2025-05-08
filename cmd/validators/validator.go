package validator

import (
	database "Synapse/cmd/validators/database"
	Print "Synapse/internal/configuration/logger/log_print"
	"fmt"
)

func ValidatorInicialize() {
	// Validando se o banco de dados está OK!
	if err := database.Init(); err != nil {
		Print.Error(err)
	} else {
		Print.Info(fmt.Sprintf("Banco de dados OK!"))
	}
}
