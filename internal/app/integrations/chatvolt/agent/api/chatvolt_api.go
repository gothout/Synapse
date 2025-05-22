package agent

import (
	agent "Synapse/internal/app/integrations/chatvolt/agent/model"
	iohelper "Synapse/internal/app/integrations/chatvolt/util/io"
	print "Synapse/internal/configuration/logger/log_print"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ChatvoltAPI interface {
	BuscarAgente(ctx context.Context, agentID, token string) (agent.Agente, error)
	EnviarMensagem(ctx context.Context, agentID, token, message, conversationId string) (agent.AgentMessageResponse, error)
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

// Manda uma mensagem para o agente da chatvolt, retorna id de conversa para manter conversacao.
func (c *chatvoltAPI) EnviarMensagem(ctx context.Context, agentID, token, message, conversationId string) (agent.AgentMessageResponse, error) {
	url := fmt.Sprintf("https://api.chatvolt.ai/agents/%s/query", agentID)

	body := map[string]interface{}{
		"query":     message,
		"streaming": false,
	}
	if conversationId != "" {
		body["conversationId"] = conversationId
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return agent.AgentMessageResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, iohelper.NopCloser(bytes.NewReader(jsonBody)))
	if err != nil {
		return agent.AgentMessageResponse{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return agent.AgentMessageResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return agent.AgentMessageResponse{}, fmt.Errorf("erro ao enviar mensagem: status %d", resp.StatusCode)
	}

	var result agent.AgentMessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return agent.AgentMessageResponse{}, err
	}

	return result, nil
}
