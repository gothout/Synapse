package agent

import model "Synapse/internal/app/integrations/chatvolt/agent/model"

// AgentConfigOutputDTO representa os dados que serão retornados após salvar a configuração do agente.
type AgentConfigResponseDTO struct {
	AgentID              string `json:"agent_id"`              // ID do agente
	Nome                 string `json:"nome"`                  // Nome do agente
	Descricao            string `json:"descricao"`             // Descrição do agente
	TokenChatVolt        string `json:"token_chatvolt"`        // Token usado na integração
	OrganizationChatVolt string `json:"organization_chatvolt"` // ID da organização na Chatvolt
}

// FromModel transforma o modelo Agente em uma resposta pronta para o cliente.
func FromModel(agente model.Agente) AgentConfigResponseDTO {
	return AgentConfigResponseDTO{
		AgentID:              agente.Id,
		Nome:                 agente.Nome,
		Descricao:            agente.Descricao,
		TokenChatVolt:        agente.TokenOrganization,
		OrganizationChatVolt: agente.OrganizationChatVoltID,
	}
}

// AgentMessageResponse representa os dados que serão retornados quando uma mensagem for enviada ao agente da Chatvolt.
type AgentMessageResponseDTO struct {
	Answer         string      `json:"answer"`
	ConversationID string      `json:"conversationId"`
	VisitorID      string      `json:"visitorId"`
	MessageID      string      `json:"messageId"`
	Sources        interface{} `json:"sources,omitempty"` // ou defina um struct específico, se quiser
	Metadata       interface{} `json:"metadata,omitempty"`
}

// FromModel transforma o modelo AgentMessageResponse em um DTO de resposta para o cliente.
func FromModelResponse(resp model.AgentMessageResponse) AgentMessageResponseDTO {
	return AgentMessageResponseDTO{
		Answer:         resp.Answer,
		ConversationID: resp.ConversationID,
		VisitorID:      resp.VisitorID,
		MessageID:      resp.MessageID,
		Sources:        resp.Sources,
		Metadata:       resp.Metadata,
	}
}

type ConfiguracaoAgentResponseDTO struct {
	ID        int64  `json:"id"`
	AgentID   string `json:"agent_id"`
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
}

func ToFormConfiguracaoAgentResponse(model model.ConfiguracaoAgent) ConfiguracaoAgentResponseDTO {
	return ConfiguracaoAgentResponseDTO{
		ID:        model.ID,
		AgentID:   model.AgentID,
		Nome:      model.Nome,
		Descricao: model.Descricao,
	}
}
