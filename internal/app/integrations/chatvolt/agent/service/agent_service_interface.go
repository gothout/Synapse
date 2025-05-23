package agent

import (
	agent "Synapse/internal/app/integrations/chatvolt/agent/model"
	"context"
)

// AgentService define as operações da regra de negócio para agentes da Chatvolt.
type AgentService interface {
	// BuscarESalvarConfiguracao busca o agente da Chatvolt e salva sua configuração
	BuscarESalvarConfiguracao(ctx context.Context, agentID string, token string) error
	BuscarConfiguracaoPorID(ctx context.Context, id int64) (agent.ConfiguracaoAgent, error)
	EnviaMensagemParaAgente(ctx context.Context, agentID int64, message string, conversationId string) (agent.AgentMessageResponse, error)
	AtualizarAgentePelaAPI(ctx context.Context, agentID int64) error
}
