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
