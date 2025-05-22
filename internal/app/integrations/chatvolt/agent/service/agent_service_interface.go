package agent

import (
	"context"
)

// AgentService define as operações da regra de negócio para agentes da Chatvolt.
type AgentService interface {
	// BuscarESalvarConfiguracao busca o agente da Chatvolt e salva sua configuração
	BuscarESalvarConfiguracao(ctx context.Context, agentID string, token string) error
}
