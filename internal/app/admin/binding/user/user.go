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

// ValidateAdminUserReadAllDTO valida o DTO de listagem paginada de usuários por empresa.
func ValidateAdminUserReadAllDTO(enterpriseID int64, page string) error {
	if enterpriseID <= 0 {
		return errors.New("enterprise_id inválido ou ausente na URI")
	}
	if strings.TrimSpace(page) == "" || page == "0" {
		return errors.New("page deve ser um número maior que zero")
	}
	return nil
}

// ValidateAdminUserReadByEmailDTO valida o DTO de busca de usuário por e-mail.
func ValidateAdminUserReadByEmailDTO(input dto.AdminUserReadByEmailDTO) error {
	if strings.TrimSpace(input.Email) == "" {
		return errors.New("e-mail ausente na URI")
	}
	if !validators.IsEmailValid(input.Email) {
		return errors.New("e-mail inválido")
	}
	return nil
}

// ValidateAdminUserUpdateDTO valida os dados recebidos para atualização de um usuário.
func ValidateAdminUserUpdateDTO(input dto.AdminUserUpdatedDTO) error {

	if input.Nome != "" && len(strings.TrimSpace(input.Nome)) < 3 {
		return errors.New("caso informado, o nome deve ter ao menos 3 caracteres")
	}

	if input.Senha != "" && len(strings.TrimSpace(input.Senha)) < 6 {
		return errors.New("caso informado, a senha deve ter ao menos 6 caracteres")
	}

	if input.Email != "" && !validators.IsEmailValid(input.Email) {
		return errors.New("caso informado, o e-mail é inválido")
	}

	return nil
}
