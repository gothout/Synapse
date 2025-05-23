package agent

// AgentConfigInputDTO representa os dados recebidos para registrar a configuração do agente Chatvolt.
type AgentConfigRequestDTO struct {
	AgentID       string `json:"agent_id" binding:"required"`       // ID do agente na Chatvolt
	TokenChatVolt string `json:"token_chatvolt" binding:"required"` // Token para autenticação na API Chatvolt
	//IntegracaoID  int    `json:"integracao_id" binding:"required"`  // ID da integração cadastrada no sistema interno
}

// AgentMessageRequest representa os dados recebidos para enviar uma mensagem para o agente da Chatvolt.
type AgentMessageRequestDTO struct {
	AgentID        int64  `json:"agent_id" binding:"required"`
	Query          string `json:"query" binding:"required"`
	ConversationID string `json:"conversationId,omitempty"`
	VisitorID      string `json:"visitorId,omitempty"`
	Streaming      bool   `json:"streaming,omitempty"` // opcional, mas pode ser fixo como false se preferir
}
