package agent

import (
	binding "Synapse/internal/app/integrations/binding/chatvolt/agent"
	dto "Synapse/internal/app/integrations/chatvolt/agent/dto"
	service "Synapse/internal/app/integrations/chatvolt/agent/service"
	"Synapse/internal/configuration/rest_err"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AgentController struct {
	service service.AgentService
}

// NewAgentController cria uma instância do controller de configuração de agentes Chatvolt
func NewAgentController(s service.AgentService) *AgentController {
	return &AgentController{service: s}
}

// PostAgentConfig godoc
// @Summary      Cadastrar configuração de agente Chatvolt
// @Description  Registra os dados de um agente da Chatvolt com base em ID, token e integração existente
// @Tags         v1 - Integração Chatvolt
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Token de integração no formato: Bearer {token}"
// @Param        request  body  dto.AgentConfigRequestDTO  true  "Dados de entrada da configuração"
// @Success      201      {object}  dto.AgentConfigResponseDTO
// @Failure      400      {object}  rest_err.RestErr
// @Failure      404      {object}  rest_err.RestErr
// @Failure      500      {object}  rest_err.RestErr
// @Router       /integrations/v1/chatvolt/agent/config [post]
func (ac *AgentController) PostAgentConfig(ctx *gin.Context) {
	var req dto.AgentConfigRequestDTO

	// Bind JSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Campos obrigatórios ausentes ou inválidos", []rest_err.Causes{
			rest_err.NewCause("body", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Validação manual via função do binding
	if err := binding.ValidateAgentConfigInputDTO(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro de validação nos dados enviados", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Chamada do serviço: busca dados e salva
	err := ac.service.BuscarESalvarConfiguracao(ctx, req.AgentID, req.TokenChatVolt)
	if err != nil {
		switch {
		case err.Error() == "agente nao encontrado":
			restErr := rest_err.NewNotFoundError("Agente não encontrado na Chatvolt")
			ctx.JSON(restErr.Code, restErr)
			return
		default:
			restErr := rest_err.NewInternalServerError("Erro ao salvar configuração do agente", []rest_err.Causes{
				rest_err.NewCause("service", err.Error()),
			})
			ctx.JSON(restErr.Code, restErr)
			return
		}
	}

	// Sucesso, sem retorno
	ctx.Status(http.StatusNoContent)
}

// PostAgentMessage godoc
// @Summary      Enviar mensagem para agente Chatvolt
// @Description  Envia uma mensagem ao agente da Chatvolt e retorna o conversationId para continuidade da conversa
// @Tags         v1 - Integração Chatvolt
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Token de integração no formato: Bearer {token}"
// @Param        request  body  dto.AgentMessageRequestDTO  true  "Dados de entrada da mensagem para o agente"
// @Success      200      {object}  dto.AgentMessageResponseDTO
// @Failure      400      {object}  rest_err.RestErr
// @Failure      404      {object}  rest_err.RestErr
// @Failure      500      {object}  rest_err.RestErr
// @Router       /integrations/v1/chatvolt/agent/message [post]
func (ac *AgentController) PostAgentMessage(ctx *gin.Context) {
	var req dto.AgentMessageRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Campos obrigatórios ausentes ou inválidos", []rest_err.Causes{
			rest_err.NewCause("body", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if err := binding.ValidateAgentMessageRequestDTO(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro de validação nos dados enviados", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Chamada do serviço retorna o model.AgentMessageResponse
	modelResp, err := ac.service.EnviaMensagemParaAgente(ctx, req.AgentID, req.Query, req.ConversationID)
	if err != nil {
		switch {
		case err.Error() == "configuração não encontrada":
			restErr := rest_err.NewNotFoundError("Configuração do agente não encontrada")
			ctx.JSON(restErr.Code, restErr)
			return
		default:
			restErr := rest_err.NewInternalServerError("Erro ao enviar mensagem ao agente", []rest_err.Causes{
				rest_err.NewCause("service", err.Error()),
			})
			ctx.JSON(restErr.Code, restErr)
			return
		}
	}

	// Transforma model em DTO e retorna
	resp := dto.FromModelResponse(modelResp)
	ctx.JSON(http.StatusOK, resp)
}

// PutAgentConfigByID godoc
// @Summary      Atualizar configuração do agente pela API da Chatvolt
// @Description  Rebusca as informações do agente na Chatvolt com base no ID e atualiza a configuração local
// @Tags         v1 - Integração Chatvolt
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Token de integração no formato: Bearer {token}"
// @Param        agent_id path int true "ID do agente registrado no sistema"
// @Success      204  "Atualizado com sucesso, sem conteúdo"
// @Failure      400  {object}  rest_err.RestErr
// @Failure      404  {object}  rest_err.RestErr
// @Failure      500  {object}  rest_err.RestErr
// @Router       /integrations/v1/chatvolt/agent/config/{agent_id} [put]
func (ac *AgentController) PutAgentConfigByID(ctx *gin.Context) {
	var req dto.PutConfiguracoesAgentRequestDTO

	// Faz o binding do path param
	if err := ctx.ShouldBindUri(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("AgentID inválido na URI", []rest_err.Causes{
			rest_err.NewCause("agent_id", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Validação manual (separada, caso queira mensagens personalizadas)
	if err := binding.ValidatePutConfiguracaoAgentRequestDTO(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Validação falhou para o agent_id", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Atualiza configurações no service.
	err := ac.service.AtualizarAgentePelaAPI(ctx, req.AgentID)
	if err != nil {
		switch {
		case err.Error() == "configuração não encontrada":
			restErr := rest_err.NewNotFoundError("Configuração do agente não encontrada")
			ctx.JSON(restErr.Code, restErr)
			return
		default:
			restErr := rest_err.NewInternalServerError("Erro ao atualizar configuração do agente", []rest_err.Causes{
				rest_err.NewCause("service", err.Error()),
			})
			ctx.JSON(restErr.Code, restErr)
			return
		}

	}

	// Sucesso, sem retorno
	ctx.Status(http.StatusNoContent)
}

// GetAgentConfigByID godoc
// @Summary      Buscar configuração do agente Chatvolt
// @Description  Retorna os dados públicos da configuração de um agente da Chatvolt por ID
// @Tags         v1 - Integração Chatvolt
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Token de integração no formato: Bearer {token}"
// @Param        agent_id path int true "ID do agente registrado no sistema"
// @Success      200  {object}  dto.AgentConfigResponseDTO
// @Failure      400  {object}  rest_err.RestErr
// @Failure      404  {object}  rest_err.RestErr
// @Failure      500  {object}  rest_err.RestErr
// @Router       /integrations/v1/chatvolt/agent/config/{agent_id} [get]
func (ac *AgentController) GetAgentConfigByID(ctx *gin.Context) {
	var req dto.GetConfiguracaoAgentRequestDTO

	// Faz o binding do path param
	if err := ctx.ShouldBindUri(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("AgentID inválido na URI", []rest_err.Causes{
			rest_err.NewCause("agent_id", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Validação manual (separada, caso queira mensagens personalizadas)
	if err := binding.ValidateGetConfiguracaoAgentRequestDTO(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Validação falhou para o agent_id", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Busca configuração no service
	config, err := ac.service.BuscarConfiguracaoPorID(ctx, req.AgentID)
	if err != nil {
		switch {
		case err.Error() == "configuração não encontrada":
			restErr := rest_err.NewNotFoundError("Configuração do agente não encontrada")
			ctx.JSON(restErr.Code, restErr)
			return
		default:
			restErr := rest_err.NewInternalServerError("Erro ao buscar configuração do agente", []rest_err.Causes{
				rest_err.NewCause("service", err.Error()),
			})
			ctx.JSON(restErr.Code, restErr)
			return
		}
	}

	// Transforma o modelo em DTO limitado e retorna
	resp := dto.ToFormConfiguracaoAgentResponse(config)
	ctx.JSON(http.StatusOK, resp)
}
