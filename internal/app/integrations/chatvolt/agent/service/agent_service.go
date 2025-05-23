package agent

import (
	integrationRepo "Synapse/internal/app/admin/integration/repository"
	contextModelMiddleware "Synapse/internal/app/admin/middleware/model"
	userRepo "Synapse/internal/app/admin/user/repository"
	api "Synapse/internal/app/integrations/chatvolt/agent/api"
	agent "Synapse/internal/app/integrations/chatvolt/agent/model"
	repository "Synapse/internal/app/integrations/chatvolt/agent/repository"
	print "Synapse/internal/configuration/logger/log_print"
	"context"
	"fmt"
	"strconv"
)

// agentService implementa a interface AgentService.
type agentService struct {
	api             api.ChatvoltAPI
	repo            repository.AgentRepository
	integrationRepo integrationRepo.Repository
	userRepo        userRepo.Repository
}

// NewAgentService cria uma nova instância do service e injeta dependências.
func NewAgentService(
	api api.ChatvoltAPI,
	repo repository.AgentRepository,
	integrationRepo integrationRepo.Repository,
	userRepo userRepo.Repository,
) AgentService {
	return &agentService{
		api:             api,
		repo:            repo,
		integrationRepo: integrationRepo,
		userRepo:        userRepo,
	}
}

// BuscarESalvarConfiguracao busca o agente da Chatvolt e salva sua configuração
// após validar se a integração informada existe no sistema.
func (s *agentService) BuscarESalvarConfiguracao(ctx context.Context, agentID string, token string) error {
	value := ctx.Value("integration")
	if value == nil {
		return fmt.Errorf("integração não encontrada no contexto")
	}

	integration, ok := value.(*contextModelMiddleware.IntegrationWithPermissions)
	fmt.Println(integration)
	if !ok {
		return fmt.Errorf("formato inválido para integração no contexto")
	}

	// Busca agente...
	agente, err := s.api.BuscarAgente(ctx, agentID, token)
	if err != nil {
		print.Error(err)
		return fmt.Errorf("%w", err)
	}

	config := map[string]interface{}{
		"agent_id":             agente.Id,
		"nome":                 agente.Nome,
		"descricao":            agente.Descricao,
		"token_chatvolt":       agente.TokenOrganization,
		"organizationChatVolt": agente.OrganizationChatVoltID,
	}

	integrationInt, err := strconv.ParseInt(integration.ID, 10, 64)
	if err != nil {
		return fmt.Errorf("ID da integração inválido: %w", err)
	}

	// Busca ID de usuario pelo token
	userID, err := s.integrationRepo.GetUserIDByToken(integration.Token)
	if err != nil {
		print.Error(err)
		return fmt.Errorf("%w", err)
	}

	// Busca ID de empresa pelo UserID
	enterpriseID, err := s.userRepo.ReadByID(userID)
	if err != nil {
		print.Error(err)
		return fmt.Errorf("%w", err)
	}

	if err := s.repo.SalvarConfiguracao(ctx, integrationInt, int64(enterpriseID.EnterpriseID), config); err != nil {
		return fmt.Errorf("erro ao salvar configuração: %w", err)
	}

	return nil
}

func (s *agentService) EnviaMensagemParaAgente(ctx context.Context, agentID int64, message string, conversationId string) (agent.AgentMessageResponse, error) {
	// Pelo ID do agente, devo buscar nas configuracoes o token da Chatvolt e id de agente, lembrando que no banco e salvo como JSON todos os dados de configuracao.

	// Busca agente...
	agente, err := s.repo.BuscarConfiguracaoPorID(ctx, agentID)
	if err != nil {
		print.Error(err)
		return agent.AgentMessageResponse{}, err
	}

	return s.api.EnviarMensagem(ctx, agente.AgentID, agente.TokenChatVolt, message, conversationId)
}

// Busca configuraçoes de agente por ID
func (s *agentService) BuscarConfiguracaoPorID(ctx context.Context, agentID int64) (agent.ConfiguracaoAgent, error) {

	agent, err := s.repo.BuscarConfiguracaoPorID(ctx, agentID)
	if err != nil {
		return agent, err
	}

	if agent.ID == 0 {
		return agent, fmt.Errorf("agente não encontrado")
	}

	return agent, nil
}

// Atualizar pela API da chatvolt
func (s *agentService) AtualizarAgentePelaAPI(ctx context.Context, agentID int64) error {

	// Buscar configurações atuais
	agent, err := s.repo.BuscarConfiguracaoPorID(ctx, agentID)
	if err != nil {
		return err
	}

	if agent.ID == 0 {
		return fmt.Errorf("agente não encontrado")
	}

	// Agora busca configurações atuais pela API da Chatvolt
	// Busca agente...
	agente, err := s.api.BuscarAgente(ctx, agent.AgentID, agent.TokenChatVolt)
	if err != nil {
		print.Error(err)
		return fmt.Errorf("%w", err)
	}

	config := map[string]interface{}{
		"agent_id":             agente.Id,
		"nome":                 agente.Nome,
		"descricao":            agente.Descricao,
		"token_chatvolt":       agente.TokenOrganization,
		"organizationChatVolt": agente.OrganizationChatVoltID,
	}

	// Atualiza agente
	if err := s.repo.AtualizarConfiguracaoPorID(ctx, agentID, config); err != nil {
		return fmt.Errorf("erro ao atualizar configuração")
	}

	return nil
}
