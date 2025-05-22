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
