package enterprise

import (
	model "Synapse/internal/app/admin/enterprise/model"
	"time"
)

type CreateEnterpriseDTO struct {
	Nome        string `json:"nome" binding:"required"`
	Cnpj        string `json:"cnpj" binding:"required"`
	Responsavel string `json:"responsavel" binding:"required"`
}

func (dto CreateEnterpriseDTO) ToModel() model.AdminEnterprise {
	return model.AdminEnterprise{
		Nome:      dto.Nome,
		Cnpj:      dto.Cnpj,
		CreatedAt: time.Now(),
	}
}
