package integration

import model "Synapse/internal/app/admin/integration/model"

// MarcaResponse representa a resposta da listagem de marcas
type MarcaResponse struct {
	ID   int64  `json:"id"`
	Nome string `json:"nome"`
}

// IntegracaoResponse representa a resposta da listagem de integrações
type IntegracaoResponse struct {
	ID      int64  `json:"id"`
	Nome    string `json:"nome"`
	MarcaID int64  `json:"marca_id"`
}

// IntegracaoDetalhadaResponse representa a resposta com join entre integração e marca
type IntegracaoDetalhadaResponse struct {
	IdIntegracao int64  `json:"id_integracao"`
	IdMarca      int64  `json:"id_marca"`
	Integracao   string `json:"integracao"`
	Empresa      string `json:"empresa"` // nome da marca
}

// FromMarcaModel converte model.Marca em DTO MarcaResponse
func FromMarcaModel(m model.Marca) MarcaResponse {
	return MarcaResponse{
		ID:   m.Id,
		Nome: m.Name,
	}
}

// FromIntegracaoModel converte model.Integration em DTO IntegracaoResponse
func FromIntegracaoModel(i model.Integration) IntegracaoResponse {
	return IntegracaoResponse{
		ID:      i.Id,
		Nome:    i.Nome,
		MarcaID: i.MarcaId,
	}
}

// FromMarcaModelList converte slice de model.Marca em slice de DTO MarcaResponse
func FromMarcaModelList(list []model.Marca) []MarcaResponse {
	resp := make([]MarcaResponse, 0, len(list))
	for _, m := range list {
		resp = append(resp, FromMarcaModel(m))
	}
	return resp
}

// FromIntegracaoModelList converte slice de model.Integration em slice de DTO IntegracaoResponse
func FromIntegracaoModelList(list []model.Integration) []IntegracaoResponse {
	resp := make([]IntegracaoResponse, 0, len(list))
	for _, i := range list {
		resp = append(resp, FromIntegracaoModel(i))
	}
	return resp
}

// FromIntegracaoDetalhadaModel converte model.IntegracaoDetalhada em DTO IntegracaoDetalhadaResponse
func FromIntegracaoDetalhadaModel(i model.IntegracaoDetalhada) IntegracaoDetalhadaResponse {
	return IntegracaoDetalhadaResponse{
		IdIntegracao: i.IdIntegracao,
		IdMarca:      i.IdMarca,
		Integracao:   i.NomeIntegracao,
		Empresa:      i.NomeMarca,
	}
}

// FromIntegracaoDetalhadaModelList converte slice de model.IntegracaoDetalhada em slice de DTO
func FromIntegracaoDetalhadaModelList(list []model.IntegracaoDetalhada) []IntegracaoDetalhadaResponse {
	resp := make([]IntegracaoDetalhadaResponse, 0, len(list))
	for _, i := range list {
		resp = append(resp, FromIntegracaoDetalhadaModel(i))
	}
	return resp
}
