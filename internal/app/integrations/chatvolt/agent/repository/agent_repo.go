package agent

import (
	agent "Synapse/internal/app/integrations/chatvolt/agent/model"
	print "Synapse/internal/configuration/logger/log_print"
	"context"
	"encoding/json"
	"fmt"

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

// BuscarConfiguracaoPorID busca as configurações de um agente armazenadas em JSON no banco, dado o ID.
func (r *agentRepo) BuscarConfiguracaoPorID(ctx context.Context, id int64) (agent.ConfiguracaoAgent, error) {
	var jsonConfig []byte
	query := `SELECT configuracoes FROM integracoes_configuracoes WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&jsonConfig)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return agent.ConfiguracaoAgent{}, fmt.Errorf("configuração não encontrada")
		}
		return agent.ConfiguracaoAgent{}, fmt.Errorf("erro ao buscar configuração no banco: %w", err)
	}

	var config agent.ConfiguracaoAgent
	if err := json.Unmarshal(jsonConfig, &config); err != nil {
		return agent.ConfiguracaoAgent{}, fmt.Errorf("erro ao fazer parse do JSON de configuração: %w", err)
	}

	config.ID = id
	return config, nil
}

// AtualizarConfiguracaoPorID atualiza o campo configuracoes (JSONB) de um agente dado o ID da linha.
func (r *agentRepo) AtualizarConfiguracaoPorID(ctx context.Context, id int64, config map[string]interface{}) error {
	jsonData, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("erro ao serializar configuração para JSON: %w", err)
	}

	query := `
		UPDATE integracoes_configuracoes
		SET configuracoes = $1,
		    atualizado_em = CURRENT_TIMESTAMP
		WHERE id = $2
	`

	tag, err := r.db.Exec(ctx, query, jsonData, id)
	if err != nil {
		print.Error(err)
		return fmt.Errorf("erro ao atualizar configuração: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("nenhuma configuração encontrada com o ID %d", id)
	}

	return nil
}
