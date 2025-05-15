package user

import (
	"errors"
	"fmt"
	"strings"

	"Synapse/internal/app/admin/pkg/validators"
	dto "Synapse/internal/app/admin/user/dto"
)

// ValidateAdminUserCreateDTO valida os dados recebidos para criação de um usuário.
func ValidateAdminUserCreateDTO(input dto.AdminUserCreateDTO) error {
	if strings.TrimSpace(input.Nome) == "" || len(input.Nome) < 3 {
		return errors.New("o nome é obrigatório e deve ter ao menos 3 caracteres")
	}

	if strings.TrimSpace(input.Email) == "" || !validators.IsEmailValid(input.Email) {
		return errors.New("e-mail inválido ou ausente")
	}

	if strings.TrimSpace(input.Senha) == "" || len(input.Senha) < 6 {
		return errors.New("a senha é obrigatória e deve ter ao menos 6 caracteres")
	}

	if strings.TrimSpace(input.Numero) == "" {
		return errors.New("o número é obrigatório")
	}

	if input.RuleID <= 0 {
		return fmt.Errorf("rule_id inválido ou ausente")
	}

	if input.EnterpriseID <= 0 {
		return fmt.Errorf("enterprise_id inválido ou ausente")
	}

	if strings.TrimSpace(input.Nome) == "" &&
		strings.TrimSpace(input.Email) == "" &&
		strings.TrimSpace(input.Senha) == "" &&
		strings.TrimSpace(input.Numero) == "" &&
		input.RuleID <= 0 &&
		input.EnterpriseID <= 0 {
		return errors.New("nenhum dado foi informado na requisição, por gentileza verificar a documentação da API")
	}

	return nil
}
