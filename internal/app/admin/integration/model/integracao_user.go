package integration

type IntegracaoUser struct {
	Id           int64 `json:"id,omitempty"`
	UserID       int64 `json:"user_id"`
	IntegracaoID int64 `json:"integracao_id"`
}
