package enterprise

import (
	dto "Synapse/internal/app/admin/enterprise/dto"
	"errors"
	"regexp"
	"strings"
)

// ValidateCreateEnterpriseDTO valida os dados recebidos para criação de uma empresa.
// Útil quando queremos ir além do binding:"required".
func ValidateCreateEnterpriseDTO(input dto.CreateEnterpriseDTO) error {
	// Remove espaços e valida campos obrigatórios
	if strings.TrimSpace(input.Nome) == "" {
		return errors.New("o nome da empresa é obrigatório")
	}
	if strings.TrimSpace(input.Cnpj) == "" {
		return errors.New("o CNPJ é obrigatório")
	}

	// Validação simples de formato de CNPJ (apenas números, 14 dígitos)
	cnpjRegex := regexp.MustCompile(`^\d{14}$`)
	if !cnpjRegex.MatchString(input.Cnpj) {
		return errors.New("o CNPJ deve conter exatamente 14 dígitos numéricos")
	}

	return nil
}
