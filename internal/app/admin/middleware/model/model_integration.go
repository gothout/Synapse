package middleware

type IntegrationWithPermissions struct {
	ID           string   `json:"id"`
	Nome         string   `json:"nome"`
	Token        string   `json:"token"`
	Permissoes   []string `json:"permissoes"`
	EnterpriseID int64    `json:"enterprise_id"`
	UserID       int64    `json:"user_id"`
}
