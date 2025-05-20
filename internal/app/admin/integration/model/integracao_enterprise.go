package integration

type IntegracaoEnterprise struct {
	Id           int64 `json:"id"`
	EnterpriseId int64 `json:"enterprise_id"`
	IntegracaoId int64 `json:"integracao_id"`
}

type IntegracaoEmpresaDetalhada struct {
	IntegracaoID int64
	Integracao   string
	Marca        string
}
