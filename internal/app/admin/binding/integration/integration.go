package integration

import (
	"errors"

	dto "Synapse/internal/app/admin/integration/dto"
)

//
// 🔹 Validação de query param (?marca_id=)
//

func ValidateGetIntegracoesByMarcaID(input dto.GetIntegracoesByMarcaIDRequest) error {
	if input.MarcaID <= 0 {
		return errors.New("marca_id inválido ou ausente na query")
	}
	return nil
}

//
// 🔹 Validação de URI param (:enterprise_id)
//

func ValidateGetIntegracoesByEnterpriseURI(input dto.GetIntegracoesByEnterpriseIDURI) error {
	if input.EnterpriseID <= 0 {
		return errors.New("enterprise_id inválido ou ausente na rota")
	}
	return nil
}

//
// 🔹 Validação de JSON para criação de vínculo entre empresa e integração
//

func ValidateCreateIntegracaoEnterprise(input dto.CreateIntegracaoEnterpriseRequest) error {
	if input.EnterpriseID <= 0 {
		return errors.New("enterprise_id inválido ou ausente no corpo da requisição")
	}
	if input.IntegracaoID <= 0 {
		return errors.New("integracao_id inválido ou ausente no corpo da requisição")
	}
	return nil
}
