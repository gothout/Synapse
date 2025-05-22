package agent

import (
	agent "Synapse/internal/app/integrations/chatvolt/agent/model"
	print "Synapse/internal/configuration/logger/log_print"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ChatvoltAPI interface {
	BuscarAgente(ctx context.Context, agentID, token string) (agent.Agente, error)
}

type chatvoltAPI struct{}

func NewChatvoltAPI() ChatvoltAPI {
	return &chatvoltAPI{}
}

// Busca agente pelo agentID da chatvolt.
func (c *chatvoltAPI) BuscarAgente(ctx context.Context, agentID, token string) (agent.Agente, error) {
	url := fmt.Sprintf("https://api.chatvolt.ai/agents/%s", agentID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return agent.Agente{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return agent.Agente{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return agent.Agente{}, fmt.Errorf("agente nao encontrado")
	}

	if resp.StatusCode != http.StatusOK {
		print.Error(err)
		return agent.Agente{}, fmt.Errorf("erro ao buscar agente: status %d", resp.StatusCode)
	}

	var agente agent.Agente
	if err := json.NewDecoder(resp.Body).Decode(&agente); err != nil {
		return agent.Agente{}, err
	}

	agente.TokenOrganization = token
	return agente, nil
}
