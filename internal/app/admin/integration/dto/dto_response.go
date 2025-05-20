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
