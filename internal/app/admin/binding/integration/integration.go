package integration

import (
	"errors"
	"strings"

	dto "Synapse/internal/app/admin/integration/dto"
	"Synapse/internal/app/admin/pkg/validators"
)

// 🔹 Validação de query param (?marca_id=)
func ValidateGetIntegracoesByMarcaID(input dto.GetIntegracoesByMarcaIDRequest) error {
	if input.MarcaID <= 0 {
		return errors.New("marca_id inválido ou ausente na query")
	}
	return nil
}

// 🔹 Validação de URI param (:enterprise_id)
func ValidateGetIntegracoesByEnterpriseURI(input dto.GetIntegracoesByEnterpriseIDURI) error {
	if input.EnterpriseID <= 0 {
		return errors.New("enterprise_id inválido ou ausente na rota")
	}
	return nil
}

// 🔹 Validação de JSON para criação de vínculo entre empresa e integração
func ValidateCreateIntegracaoEnterprise(input dto.CreateIntegracaoEnterpriseRequest) error {
	if input.EnterpriseID <= 0 {
		return errors.New("enterprise_id inválido ou ausente no corpo da requisição")
	}
	if input.IntegracaoID <= 0 {
		return errors.New("integracao_id inválido ou ausente no corpo da requisição")
	}
	return nil
}

// 🔹 Validação de JSON para remover vínculo entre empresa e integração
func ValidateDeleteIntegracaoEnterprise(input dto.DeleteIntegracaoEnterpriseRequest) error {
	if input.EnterpriseID <= 0 {
		return errors.New("enterprise_id inválido ou ausente")
	}
	if input.IntegracaoID <= 0 {
		return errors.New("integracao_id inválido ou ausente")
	}
	return nil
}

// Valida o JSON para criar vinculo entre empresa e integração
func ValidateCreateIntegracaoUser(input dto.CreateIntegracaoUserRequest) error {
	if input.UserID <= 0 {
		return errors.New("user_id inválido ou ausente")
	}
	if input.IntegracaoID <= 0 {
		return errors.New("integracao_id inválido ou ausente")
	}
	return nil
}

// Valida CreateTokenIntegracaoRequest para criar um token de integração
func ValidateCreateTokenIntegracao(input dto.CreateTokenIntegracaoRequest) error {
	if strings.TrimSpace(input.Email) == "" {
		return errors.New("email é obrigatório")
	}
	if !validators.IsEmailValid(input.Email) {
		return errors.New("email inválido")
	}
	if len(strings.TrimSpace(input.Senha)) < 6 {
		return errors.New("a senha deve conter pelo menos 6 caracteres")
	}
	if input.IntegracaoID <= 0 {
		return errors.New("integracao_id inválido ou ausente")
	}
	return nil
}
