package agent

// AgentConfigInputDTO representa os dados recebidos para registrar a configuração do agente Chatvolt.
type AgentConfigRequestDTO struct {
	AgentID       string `json:"agent_id" binding:"required"`       // ID do agente na Chatvolt
	TokenChatVolt string `json:"token_chatvolt" binding:"required"` // Token para autenticação na API Chatvolt
	//IntegracaoID  int    `json:"integracao_id" binding:"required"`  // ID da integração cadastrada no sistema interno
}
