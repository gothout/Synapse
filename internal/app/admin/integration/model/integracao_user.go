package integration

type IntegracaoUser struct {
	Id           int64 `json:"id"`
	UserId       int64 `json:"user_id"`
	IntegracaoId int64 `json:"integracao_id"`
}
