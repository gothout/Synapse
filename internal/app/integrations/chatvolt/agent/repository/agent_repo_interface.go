package agent

import (
	"context"
)

type AgentRepository interface {
	// SalvarConfiguracao insere ou atualiza a configuração do agente no campo JSONB da tabela.
	SalvarConfiguracao(ctx context.Context, integracaoID int64, config map[string]interface{}) error
}
