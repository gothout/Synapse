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

func (dto CreateEnterpriseDTO) ToModel() model.AdminEnterprise {
	return model.AdminEnterprise{
		Nome:      dto.Nome,
		Cnpj:      dto.Cnpj,
		CreatedAt: time.Now(),
	}
}
