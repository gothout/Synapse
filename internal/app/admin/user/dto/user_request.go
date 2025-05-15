package user

import model "Synapse/internal/app/admin/user/model"

// AdminUserCreateDTO representa a estrutura de entrada para criar novo usuario.
type AdminUserCreateDTO struct {
	Nome         string `json:"nome" binding:"required,min=3"`
	Email        string `json:"email" binding:"required,email"`
	Senha        string `json:"senha" binding:"required,min=6"`
	Numero       string `json:"numero" binding:"required"`
	RuleID       int64  `json:"rule_id" binding:"required"`
	EnterpriseID int64  `json:"enterprise_id" binding:"required"`
}

// ToModel converte um DTO de criação para um model.Admin
func (dto AdminUserCreateDTO) ToModel() model.User {
	return model.User{
		Nome:         dto.Nome,
		Email:        dto.Email,
		Senha:        dto.Senha,
		Numero:       dto.Numero,
		RuleID:       dto.RuleID,
		EnterpriseID: dto.EnterpriseID,
	}
}
