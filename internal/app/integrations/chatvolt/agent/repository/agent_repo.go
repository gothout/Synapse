package agent

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgxpool"
)

// agentRepo implementa o AgentRepository e lida com o banco de dados.
type agentRepo struct {
	db *pgxpool.Pool
}

// NewAgentRepository retorna uma nova instância do repositório.
func NewAgentRepository(db *pgxpool.Pool) AgentRepository {
	return &agentRepo{db}
}

// SalvarConfiguracao insere ou atualiza a configuração do agente no campo JSONB da tabela.
func (r *agentRepo) SalvarConfiguracao(ctx context.Context, integracaoID int64, enterprise_id int64, config map[string]interface{}) error {
	jsonData, err := json.Marshal(config)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO integracoes_configuracoes (integracao_id,enterprise_id, configuracoes)
		VALUES ($1, $2, $3)
		ON CONFLICT (integracao_id)
		DO UPDATE SET configuracoes = EXCLUDED.configuracoes, atualizado_em = CURRENT_TIMESTAMP
	`

	_, err = r.db.Exec(ctx, query, integracaoID, enterprise_id, jsonData)
	return err
}
