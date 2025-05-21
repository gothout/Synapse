package middleware

type IntegrationWithPermissions struct {
	ID         string
	Nome       string
	Token      string
	Permissoes []string
}
