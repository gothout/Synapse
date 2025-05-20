package integration

type Integration struct {
	Id      int64  `json:"id"`
	Nome    string `json:"nome"`
	MarcaId int64  `json:"marcaId"`
}

type IntegracaoDetalhada struct {
	IdIntegracao   int64
	IdMarca        int64
	NomeIntegracao string
	NomeMarca      string
}
