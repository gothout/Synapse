package user

import model "Synapse/internal/app/admin/user/model"

// AdminUserTokenResponseDTO representa o retorno de login
type AdminUserTokenResponseDTO struct {
	ID           int64  `json:"id"`
	Nome         string `json:"nome"`
	Email        string `json:"email"`
	Numero       string `json:"numero"`
	EnterpriseID int64  `json:"enterprise_id"`
	RuleID       int64  `json:"rule_id"`
	Token        string `json:"token"`
}

// UserResponseDTO representa os dados retornados de um usuário
type UserResponseDTO struct {
	ID           int64  `json:"id"`
	Nome         string `json:"nome"`
	Email        string `json:"email"`
	Numero       string `json:"numero"`
	RuleID       int64  `json:"rule_id"`
	EnterpriseID int64  `json:"enterprise_id"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

// FromModel converte um model.User para DTO de resposta
func FromModel(u model.User) UserResponseDTO {
	return UserResponseDTO{
		ID:           u.ID,
		Nome:         u.Nome,
		Email:        u.Email,
		Numero:       u.Numero,
		RuleID:       u.RuleID,
		EnterpriseID: u.EnterpriseID,
		CreatedAt:    u.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    u.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// FromModel updated converte um model.User para DTO de response.
func FromModelUpdated(u model.User) UserResponseDTO {
	return UserResponseDTO{
		ID:           u.ID,
		Nome:         u.Nome,
		Email:        u.Email,
		Numero:       u.Numero,
		RuleID:       u.RuleID,
		EnterpriseID: u.EnterpriseID,
		UpdatedAt:    u.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func FromModelToken(user model.User, token string) AdminUserTokenResponseDTO {
	return AdminUserTokenResponseDTO{
		ID:           user.ID,
		Nome:         user.Nome,
		Email:        user.Email,
		Numero:       user.Numero,
		EnterpriseID: user.EnterpriseID,
		RuleID:       user.RuleID,
		Token:        token,
	}
}

// FromModelList converte uma lista de usuários para uma lista de DTOs
func FromModelList(users []model.User) []UserResponseDTO {
	response := make([]UserResponseDTO, 0, len(users))
	for _, u := range users {
		response = append(response, FromModel(u))
	}
	return response
}
