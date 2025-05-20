package integration

import model "Synapse/internal/app/admin/integration/model"

//
// 🔹 Marca
//

// MarcaResponse representa a resposta da listagem de marcas
type MarcaResponse struct {
	ID   int64  `json:"id"`
	Nome string `json:"nome"`
}

func FromMarcaModel(m model.Marca) MarcaResponse {
	return MarcaResponse{
		ID:   m.Id,
		Nome: m.Name,
	}
}

func FromMarcaModelList(list []model.Marca) []MarcaResponse {
	resp := make([]MarcaResponse, 0, len(list))
	for _, m := range list {
		resp = append(resp, FromMarcaModel(m))
	}
	return resp
}

//
// 🔹 Integração simples
//

// IntegracaoResponse representa a resposta da listagem de integrações
type IntegracaoResponse struct {
	ID      int64  `json:"id"`
	Nome    string `json:"nome"`
	MarcaID int64  `json:"marca_id"`
}

func FromIntegracaoModel(i model.Integration) IntegracaoResponse {
	return IntegracaoResponse{
		ID:      i.Id,
		Nome:    i.Nome,
		MarcaID: i.MarcaId,
	}
}

func FromIntegracaoModelList(list []model.Integration) []IntegracaoResponse {
	resp := make([]IntegracaoResponse, 0, len(list))
	for _, i := range list {
		resp = append(resp, FromIntegracaoModel(i))
	}
	return resp
}

//
// 🔹 Integração detalhada (JOIN com marca)
//

type IntegracaoDetalhadaResponse struct {
	IdIntegracao int64  `json:"id_integracao"`
	IdMarca      int64  `json:"id_marca"`
	Integracao   string `json:"integracao"`
	Empresa      string `json:"empresa"` // nome da marca
}

func FromIntegracaoDetalhadaModel(i model.IntegracaoDetalhada) IntegracaoDetalhadaResponse {
	return IntegracaoDetalhadaResponse{
		IdIntegracao: i.IdIntegracao,
		IdMarca:      i.IdMarca,
		Integracao:   i.NomeIntegracao,
		Empresa:      i.NomeMarca,
	}
}

func FromIntegracaoDetalhadaModelList(list []model.IntegracaoDetalhada) []IntegracaoDetalhadaResponse {
	resp := make([]IntegracaoDetalhadaResponse, 0, len(list))
	for _, i := range list {
		resp = append(resp, FromIntegracaoDetalhadaModel(i))
	}
	return resp
}

//
// 🔹 Integrações liberadas por empresa
//

type IntegracaoEmpresaDetalhadaResponse struct {
	IntegracaoID int64  `json:"integracao_id"`
	Nome         string `json:"nome"`
	Marca        string `json:"marca"`
}

func FromIntegracaoEmpresaDetalhadaModel(m model.IntegracaoEmpresaDetalhada) IntegracaoEmpresaDetalhadaResponse {
	return IntegracaoEmpresaDetalhadaResponse{
		IntegracaoID: m.IntegracaoID,
		Nome:         m.Integracao,
		Marca:        m.Marca,
	}
}

func FromIntegracaoEmpresaDetalhadaModelList(list []model.IntegracaoEmpresaDetalhada) []IntegracaoEmpresaDetalhadaResponse {
	resp := make([]IntegracaoEmpresaDetalhadaResponse, 0, len(list))
	for _, i := range list {
		resp = append(resp, FromIntegracaoEmpresaDetalhadaModel(i))
	}
	return resp
}

// IntegracaoUsuarioResponse representa a resposta da listagem de permissões do usuário
type IntegracaoUsuarioResponse struct {
	ID    int64  `json:"id"`
	Nome  string `json:"nome"`
	Marca string `json:"marca"`
}

func FromIntegracaoUsuarioModelList(list []model.IntegracaoUsuarioDetalhada) []IntegracaoUsuarioResponse {
	resp := make([]IntegracaoUsuarioResponse, 0, len(list))
	for _, i := range list {
		resp = append(resp, IntegracaoUsuarioResponse{
			ID:    i.ID,
			Nome:  i.Nome,
			Marca: i.Marca,
		})
	}
	return resp
}
