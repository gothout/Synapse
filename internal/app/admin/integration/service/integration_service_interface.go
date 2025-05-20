package integration

import integration "Synapse/internal/app/admin/integration/model"

type Service interface {
	// Retorna todas as integrações
	GetAllIntegrations() ([]integration.Integration, error)
	//Busca todas as marcas da integração
	GetAllMarcas() ([]integration.Marca, error)
	//Buscar todas as integrações baseado no MarcaID
	GetIntegrationsByMarcaID(marcaID int64) ([]integration.IntegracaoDetalhada, error)
	// Vincular empresa a integração
	CreateIntegracaoEnterprise(data integration.IntegracaoEnterprise) error
	// Retorna detalhe de integrações para X empresa
	GetIntegracoesByEnterpriseID(enterpriseID int64) ([]integration.IntegracaoEmpresaDetalhada, error)
}
