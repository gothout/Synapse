package enterprise

import (
	dto "Synapse/internal/app/admin/enterprise/dto"
	validators "Synapse/internal/app/admin/pkg/validators"
	"errors"
	"fmt"
	"strings"
)

// ValidateCreateEnterpriseDTO valida os dados recebidos para criação de uma empresa.
func ValidateCreateEnterpriseDTO(input dto.CreateEnterpriseDTO) error {
	if strings.TrimSpace(input.Nome) == "" {
		return errors.New("o nome da empresa é obrigatório")
	}
	if strings.TrimSpace(input.Cnpj) == "" {
		return errors.New("o CNPJ é obrigatório")
	}

	// Validação completa de CNPJ (estrutura e dígitos verificadores)
	if err := validators.ValidateCNPJ(input.Cnpj); err != nil {
		return err
	}

	return nil
}

// ValidateReadEnterpriseCNPJDTO valida o DTO de leitura por CNPJ (rota /empresa/:cnpj)
func ValidateReadEnterpriseCNPJDTO(input dto.ReadEnterpriseCNPJDTO) error {
	if strings.TrimSpace(input.Cnpj) == "" {
		return errors.New("o CNPJ é obrigatório")
	}

	if err := validators.ValidateCNPJ(input.Cnpj); err != nil {
		return err
	}

	return nil
}

// ValidateReadEnterpriseNOMEDTO valida o DTO de leitura por NOME (rota /empresa/:cnpj)
func ValidateReadEnterpriseNOMEDTO(input dto.ReadEnterpriseNOMEDTO) error {
	if strings.TrimSpace(input.Nome) == "" {
		return errors.New("o NOME é obrigatório")
	}
	return nil
}

// ValidateReadEnterpriseNOMEDTO valida o DTO de leitura por NOME (rota /empresa/:cnpj)
func ValidateReadEnterpriseIDDTO(input dto.ReadEnterpriseIDDTO) error {
	if input.ID <= 0 {
		return errors.New("o ID é obrigatório e deve ser maior que zero")
	}
	return nil
}

func ValidateReadEnterpriseAllDTO(input dto.ReadEnterpriseAllDTO) (int, error) {
	page := input.Page

	// Se não vier nenhum valor no query param, Gin deixa como zero
	if page == 0 {
		page = -1 // valor padrão se não enviado
	}

	if page < -1 {
		return 0, fmt.Errorf("o valor de 'page' não pode ser menor que -1")
	}

	return page, nil
}

func ValidateUpdateEnterpriseDTO(input dto.UpdateEnterpriseByCNPJDTO) error {
	// Valida o CNPJ original da URI
	cnpjOriginal := strings.TrimSpace(input.Cnpj)
	if cnpjOriginal == "" {
		return errors.New("o CNPJ de identificação da empresa é obrigatório")
	}
	if err := validators.ValidateCNPJ(cnpjOriginal); err != nil {
		return fmt.Errorf("CNPJ original inválido: %w", err)
	}

	// Verifica se ao menos um campo para atualização foi informado
	nome := strings.TrimSpace(input.Nome)
	newCNPJ := strings.TrimSpace(input.NewCNPJ)

	if nome == "" && newCNPJ == "" {
		return errors.New("é necessário informar ao menos um dado para atualização (nome ou novo CNPJ)")
	}

	// Valida o novo nome, se informado
	if nome != "" && len(nome) < 2 {
		return errors.New("o nome informado é muito curto")
	}

	// Valida o novo CNPJ, se informado
	if newCNPJ != "" {
		if err := validators.ValidateCNPJ(newCNPJ); err != nil {
			return fmt.Errorf("novo CNPJ inválido: %w", err)
		}
	}

	return nil
}

// ValidateDeleteEnterpriseByCNPJDTO valida o DTO de exclusão por CNPJ (rota /empresa/cnpj/{cnpj})
func ValidateDeleteEnterpriseByCNPJDTO(input dto.DeleteEnterpriseByCNPJDTO) error {
	if strings.TrimSpace(input.Cnpj) == "" {
		return errors.New("o CNPJ é obrigatório")
	}

	if err := validators.ValidateCNPJ(input.Cnpj); err != nil {
		return fmt.Errorf("CNPJ inválido: %w", err)
	}

	return nil
}
