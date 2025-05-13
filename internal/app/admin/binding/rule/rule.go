package rule

import (
	dto "Synapse/internal/app/admin/rule/dto"
	"errors"
	"fmt"
)

// ValidateGetRulesQueryDTO valida os parâmetros de query de paginação
func ValidateGetRulesQueryDTO(input dto.GetRulesQueryDTO) (int, int, error) {
	if input.Limit <= 0 {
		return 0, 0, errors.New("o parâmetro 'limit' deve ser maior que zero")
	}
	if input.Offset < 0 {
		return 0, 0, errors.New("o parâmetro 'offset' não pode ser negativo")
	}

	return input.Limit, input.Offset, nil
}

// ValidateReadRuleByIDDTO valida o DTO de leitura de uma regra por ID
func ValidateReadRuleByIDDTO(id int64) error {
	if id <= 0 {
		return fmt.Errorf("o ID da regra deve ser maior que zero")
	}
	return nil
}

// ValidateReadPermissionsByRuleIDDTO valida o ID da regra antes de buscar permissões
func ValidateReadPermissionsByRuleIDDTO(ruleID int64) error {
	if ruleID <= 0 {
		return fmt.Errorf("o ID da regra é obrigatório e deve ser maior que zero")
	}
	return nil
}
