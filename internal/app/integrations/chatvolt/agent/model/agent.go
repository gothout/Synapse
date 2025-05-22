package agent

// Agente representa os campos filtrados do retorno da API da Chatvolt
type Agente struct {
	Id                     string `json:"id"`
	Nome                   string `json:"name"`           // "name" → Nome
	Descricao              string `json:"description"`    // "description" → Descricao
	OrganizationChatVoltID string `json:"organizationId"` // "organizationId" → OrganizationChatVoltID
	TokenOrganization      string `json:"-"`              // será injetado manualmente, não vem no JSON
}
