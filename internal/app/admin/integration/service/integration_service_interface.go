package integration

import integration "Synapse/internal/app/admin/integration/model"

type Service interface {
	// Retorna todas as integrações
	GetAllIntegrations() ([]integration.Integration, error)
	GetAllMarcas() ([]integration.Marca, error)
}
