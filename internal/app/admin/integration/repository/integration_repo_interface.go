package integration

import integration "Synapse/internal/app/admin/integration/model"

type Repository interface {
	// Busca todas as integrações
	GetAllIntegrations() ([]integration.Integration, error)
	//Busca todas as marcas da integração
	GetAllMarcas() ([]integration.Marca, error)
}
