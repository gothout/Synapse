package user

import (
	model "Synapse/internal/app/admin/user/model"
	"strings"
)

// AdminUserCreateDTO representa a estrutura de entrada para criar novo usuario.
type AdminUserCreateDTO struct {
	Nome         string `json:"nome" binding:"required,min=3"`
	Email        string `json:"email" binding:"required,email"`
	Senha        string `json:"senha" binding:"required,min=6"`
	Numero       string `json:"numero" binding:"required"`
	RuleID       int64  `json:"rule_id" binding:"required"`
	EnterpriseID int64  `json:"enterprise_id" binding:"required"`
}

// AdminUserUpdatedDTO representa a estrutura de entrada para atualizar um novo usuario.
type AdminUserUpdatedDTO struct {
	Nome         string `json:"nome"`
	Email        string `json:"email"`
	Senha        string `json:"senha"`
	Numero       string `json:"numero"`
	RuleID       int64  `json:"rule_id"`
	EnterpriseID int64  `json:"enterprise_id"`
}

// AdminUpdatedUserDTO representa o userID para atualizar.
type AdminUpdatedUserDTO struct {
	UserID int64 `uri:"id" binding:"required"`
}

// AdminUserReadAllDTO representa o enterprise ID para validação.
type AdminUserReadAllURI struct {
	EnterpriseID int64 `uri:"enterprise_id" binding:"required"`
}
type AdminUserReadAllQuery struct {
	Page string `form:"page" binding:"required"`
}

// AdminUserReadByEmailDTO representa o email para validação.
type AdminUserReadByEmailDTO struct {
	Email string `uri:"email" binding:"required"`
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

// ToModelUpdated converte o DTO de atualização para model.User
func (dto AdminUserUpdatedDTO) ToModelUpdated() model.User {
	return model.User{
		Nome:         strings.TrimSpace(dto.Nome),
		Email:        strings.TrimSpace(dto.Email),
		Senha:        strings.TrimSpace(dto.Senha),
		Numero:       strings.TrimSpace(dto.Numero),
		RuleID:       dto.RuleID,
		EnterpriseID: dto.EnterpriseID,
	}
}
