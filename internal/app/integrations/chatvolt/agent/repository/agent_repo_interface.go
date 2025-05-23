package agent

import (
	agent "Synapse/internal/app/integrations/chatvolt/agent/model"
	"context"
)

type AgentRepository interface {
	// SalvarConfiguracao insere ou atualiza a configuração do agente no campo JSONB da tabela.
	SalvarConfiguracao(ctx context.Context, integracaoID int64, enterprise_id int64, config map[string]interface{}) error
	// Busca configuracoes de agente por ID do banco.
	BuscarConfiguracaoPorID(ctx context.Context, id int64) (agent.ConfiguracaoAgent, error)
	// Atualiza configurações de agente por ID
	AtualizarConfiguracaoPorID(ctx context.Context, id int64, config map[string]interface{}) error
	// Listar agentes por empresaID
	BuscarAgentesPorEmpresaID(ctx context.Context, empresaID int64) ([]agent.ConfiguracaoAgent, error)
	// DeleteConfigByID
	DeleteConfigByID(ctx context.Context, id int64, empresaId int64) error
	// Busca todas as configurações pelo ID da configuração de integração
	BuscaEmpresaDeAgenteByAgentId(ctx context.Context, agentId int64) (int64, error)
}
