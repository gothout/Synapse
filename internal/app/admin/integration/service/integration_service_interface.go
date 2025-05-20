package integration

import integration "Synapse/internal/app/admin/integration/model"

type Service interface {
	// Retorna todas as integrações
	GetAllIntegrations() ([]integration.Integration, error)
	//Busca todas as marcas da integração
	GetAllMarcas() ([]integration.Marca, error)
	//Buscar todas as integrações baseado no MarcaID
	GetIntegrationsByMarcaID(marcaID int64) ([]integration.IntegracaoDetalhada, error)
}
