package rule

import model "Synapse/internal/app/admin/rule/model"

// AdminRuleCreateDTO representa os dados para criação de uma nova regra
type AdminRuleCreateDTO struct {
	Name string `json:"name" binding:"required,min=3"`
}

// GetRulesQueryDTO representa os parâmetros de query para paginação
type GetRulesQueryDTO struct {
	Limit  int `form:"limit" binding:"required,min=1"`
	Offset int `form:"offset" binding:"min=0"`
}

// ToModel converte um DTO de criação para um model.AdminRule
func (dto AdminRuleCreateDTO) ToModel() model.AdminRule {
	return model.AdminRule{
		Name: dto.Name,
	}
}
