package integration

type IntegracaoUser struct {
	Id           int64 `json:"id,omitempty"`
	UserID       int64 `json:"user_id"`
	IntegracaoID int64 `json:"integracao_id"`
}

// IntegracaoUsuarioDetalhada representa os dados de uma integração que o usuário possui
type IntegracaoUsuarioDetalhada struct {
	ID    int64  `json:"id"`
	Nome  string `json:"nome"`
	Marca string `json:"marca"`
}
