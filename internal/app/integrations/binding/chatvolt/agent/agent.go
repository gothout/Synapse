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
