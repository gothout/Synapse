package integration

import model "Synapse/internal/app/admin/integration/model"

//
// 🔹 Request: Filtro por marca_id (query/json)
//

// GetIntegracoesByMarcaIDRequest representa o filtro por marca_id
// Usado para buscas via query param (?marca_id=) ou JSON se necessário
type GetIntegracoesByMarcaIDRequest struct {
	MarcaID int64 `form:"marca_id" json:"marca_id" binding:"required"`
}

// ToModelMarca converte o DTO para o model.Marca
func (req GetIntegracoesByMarcaIDRequest) ToModelMarca() model.Marca {
	return model.Marca{
		Id: req.MarcaID,
	}
}

//
// 🔹 Request: URI param para enterprise_id
//

// GetIntegracoesByEnterpriseIDURI representa o parâmetro de rota :enterprise_id
// Exemplo de rota: /integration/enterprise/:enterprise_id
type GetIntegracoesByEnterpriseIDURI struct {
	EnterpriseID int64 `uri:"enterprise_id" binding:"required"`
}

// ToID retorna apenas o ID como inteiro (semântica de simplificação)
func (uri GetIntegracoesByEnterpriseIDURI) ToID() int64 {
	return uri.EnterpriseID
}

//
// 🔹 Request: corpo para criação de vínculo entre empresa e integração
//

// CreateIntegracaoEnterpriseRequest representa os dados do JSON enviados para criar o vínculo
// entre uma empresa e uma integração.
type CreateIntegracaoEnterpriseRequest struct {
	EnterpriseID int64 `json:"enterprise_id" binding:"required"`
	IntegracaoID int64 `json:"integracao_id" binding:"required"`
}

// ToModelIntegracaoEnterprise converte o DTO para o model.IntegracaoEnterprise
func (req CreateIntegracaoEnterpriseRequest) ToModelIntegracaoEnterprise() model.IntegracaoEnterprise {
	return model.IntegracaoEnterprise{
		EnterpriseId: req.EnterpriseID,
		IntegracaoId: req.IntegracaoID,
	}
}
