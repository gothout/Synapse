package integration

import (
	"errors"

	dto "Synapse/internal/app/admin/integration/dto"
)

// ValidateGetIntegracoesByMarcaID valida se o ID da marca foi informado corretamente.
func ValidateGetIntegracoesByMarcaID(input dto.GetIntegracoesByMarcaIDRequest) error {
	if input.MarcaID <= 0 {
		return errors.New("marca_id inválido ou ausente na query")
	}
	return nil
}
