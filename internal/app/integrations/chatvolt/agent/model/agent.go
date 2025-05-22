package agent

// Agente representa os campos filtrados do retorno da API da Chatvolt
type Agente struct {
	Id                     string `json:"id"`
	Nome                   string `json:"name"`           // "name" → Nome
	Descricao              string `json:"description"`    // "description" → Descricao
	OrganizationChatVoltID string `json:"organizationId"` // "organizationId" → OrganizationChatVoltID
	TokenOrganization      string `json:"-"`              // será injetado manualmente, não vem no JSON
}

type ConfiguracaoAgent struct {
	ID                   int64  `json:"id"`
	AgentID              string `json:"agent_id"`
	Nome                 string `json:"nome"`
	Descricao            string `json:"descricao"`
	TokenChatVolt        string `json:"token_chatvolt"`
	OrganizationChatVolt string `json:"organizationChatVolt"`
}

type AgentMessageResponse struct {
	Answer         string
	ConversationID string
	VisitorID      string
	MessageID      string
	Sources        interface{}
	Metadata       interface{}
}
