package agent

import (
	dto "Synapse/internal/app/integrations/chatvolt/agent/dto"
	"errors"
	"strings"
)

// ValidateAgentConfigInputDTO valida os dados recebidos no DTO de criação de configuração do agente Chatvolt.
func ValidateAgentConfigInputDTO(input dto.AgentConfigRequestDTO) error {
	if strings.TrimSpace(input.AgentID) == "" {
		return errors.New("o AgentID é obrigatório")
	}

	if len(input.AgentID) < 6 {
		return errors.New("o AgentID é muito curto")
	}

	if strings.TrimSpace(input.TokenChatVolt) == "" {
		return errors.New("o token da Chatvolt é obrigatório")
	}

	if len(input.TokenChatVolt) < 10 {
		return errors.New("o token da Chatvolt parece inválido (muito curto)")
	}

	return nil
}

// ValidateAgentMessageRequestDTO valida os dados recebidos para envio de mensagem ao agente.
func ValidateAgentMessageRequestDTO(input dto.AgentMessageRequestDTO) error {
	if input.AgentID <= 0 {
		return errors.New("o ID do agente deve ser maior que zero")
	}

	if strings.TrimSpace(input.Query) == "" {
		return errors.New("a mensagem (query) é obrigatória")
	}

	if len(input.Query) < 3 {
		return errors.New("a mensagem é muito curta")
	}

	if input.ConversationID != "" && len(input.ConversationID) < 10 {
		return errors.New("o conversationId parece inválido")
	}

	if input.VisitorID != "" && len(input.VisitorID) < 5 {
		return errors.New("o visitorId parece inválido")
	}

	return nil
}

// ValidateGetConfiguracaoAgentRequestDTO valida os dados recebidos para buscar configurações do agente.
func ValidateGetConfiguracaoAgentRequestDTO(input dto.GetConfiguracaoAgentRequestDTO) error {
	if input.AgentID <= 0 {
		return errors.New("o agent_id deve ser maior que zero")
	}
	return nil
}

// ValidatePutConfiguracaoAgentRequestDTO
func ValidatePutConfiguracaoAgentRequestDTO(input dto.PutConfiguracoesAgentRequestDTO) error {
	if input.AgentID <= 0 {
		return errors.New("o agent_id deve ser maior que zero")
	}
	return nil
}
