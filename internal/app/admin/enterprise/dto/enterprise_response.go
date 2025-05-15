package enterprise

import (
	model "Synapse/internal/app/admin/enterprise/model"
	"time"
)

// EnterpriseResponseDTO representa a estrutura de resposta da empresa.
type EnterpriseResponseDTO struct {
	ID        int64     `json:"id"`
	Nome      string    `json:"nome"`
	Cnpj      string    `json:"cnpj"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
type EnterpriseUpdatedResponseDTO struct {
	Nome      string    `json:"nome"`
	Cnpj      string    `json:"cnpj"`
	UpdatedAt time.Time `json:"updated_at"`
}
type EnterpriseDeletedResponseDTO struct {
	Cnpj string `json:"cnpj"`
}

// FromModel converte um model.AdminEnterprise para um EnterpriseResponseDTO.
func FromModel(ent model.AdminEnterprise) EnterpriseResponseDTO {
	return EnterpriseResponseDTO{
		ID:        ent.ID,
		Nome:      ent.Nome,
		Cnpj:      ent.Cnpj,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}
}

func FromModelUpdateResponse(ent model.AdminEnterprise) EnterpriseUpdatedResponseDTO {
	return EnterpriseUpdatedResponseDTO{
		Nome:      ent.Nome,
		Cnpj:      ent.Cnpj,
		UpdatedAt: ent.UpdatedAt,
	}
}

// FromModelList converte uma lista de model.AdminEnterprise para []EnterpriseResponseDTO.
func FromModelList(ents []model.AdminEnterprise) []EnterpriseResponseDTO {
	response := make([]EnterpriseResponseDTO, 0, len(ents))
	for _, ent := range ents {
		response = append(response, FromModel(ent))
	}
	return response
}
