package user

type User struct {
	ID           int64  `json:"id"`
	Nome         string `json:"nome"`
	Email        string `json:"email"`
	Senha        string `json:"senha"`
	Numero       string `json:"numero"`
	Token        string `json:"token"`
	RuleID       int64  `json:"rule_id"`
	EnterpriseID int64  `json:"enterprise_id"`
}
