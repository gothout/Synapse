package enterprise

import (
	model "Synapse/internal/app/admin/enterprise/model"
	"time"
)

type CreateEnterpriseDTO struct {
	Nome string `json:"nome" binding:"required"`
	Cnpj string `json:"cnpj" binding:"required"`
}

type ReadEnterpriseCNPJDTO struct {
	Cnpj string `uri:"cnpj" binding:"required"`
}
type ReadEnterpriseIDDTO struct {
	ID int64 `uri:"id" binding:"required"`
}

type ReadEnterpriseNOMEDTO struct {
	Nome string `uri:"nome" binding:"required"`
}

type ReadEnterpriseAllDTO struct {
	Page int `form:"page"`
}

type UpdateEnterpriseByCNPJDTO struct {
	Cnpj    string `uri:"cnpj" binding:"required"` // CNPJ original (na URL)
	Nome    string `json:"nome"`                   // nome opcional
	NewCNPJ string `json:"cnpj"`                   // novo CNPJ opcional
}

func (dto CreateEnterpriseDTO) ToModel() model.AdminEnterprise {
	return model.AdminEnterprise{
		Nome:      dto.Nome,
		Cnpj:      dto.Cnpj,
		CreatedAt: time.Now(),
	}
}

func (dto UpdateEnterpriseByCNPJDTO) UpdateToModel() model.AdminEnterprise {
	return model.AdminEnterprise{
		Nome:      dto.Nome,
		Cnpj:      dto.NewCNPJ,
		UpdatedAt: time.Now(),
	}
}
