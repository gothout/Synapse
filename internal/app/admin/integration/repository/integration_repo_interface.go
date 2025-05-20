package integration

import (
	integration "Synapse/internal/app/admin/integration/model"
	"context"
)

type Repository interface {
	// Busca todas as integrações
	GetAllIntegrations() ([]integration.Integration, error)
	//Busca todas as marcas da integração
	GetAllMarcas() ([]integration.Marca, error)
	//Buscar todas as integrações baseado no MarcaID
	GetIntegracoesDetalhadasByMarcaID(marcaID int64) ([]integration.IntegracaoDetalhada, error)

	// Vincular empresa a integração
	CreateIntegracaoEnterprise(data integration.IntegracaoEnterprise) error
	// Busca integracao por ID
	GetIntegracaoByID(ctx context.Context, id int64) (*integration.Integration, error)
	// Retorna detalhe de integrações para X empresa
	GetIntegracoesByEnterpriseID(enterpriseID int64) ([]integration.IntegracaoEmpresaDetalhada, error)
	// Remover integração de empresa
	DeleteIntegracaoFromEnterprise(enterpriseID, integrationID int64) error
	// Vincular usuario a integração
	CreateIntegracaoUser(data integration.IntegracaoUser) error
	// Criar token para acessar integração
	SaveIntegracaoToken(userID, integracaoID int64, token string) error
	CheckUserHasIntegracao(userID, integracaoID int64) (bool, error)
}
