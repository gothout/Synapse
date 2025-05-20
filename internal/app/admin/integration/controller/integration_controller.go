package controller

import (
	dto "Synapse/internal/app/admin/integration/dto"
	integrationService "Synapse/internal/app/admin/integration/service"
	"Synapse/internal/configuration/rest_err"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IntegrationController struct {
	service integrationService.Service
}

func NewIntegrationController(s integrationService.Service) *IntegrationController {
	return &IntegrationController{service: s}
}

// GetAllIntegrations godoc
// @Summary      Listar integrações
// @Description  Retorna todas as integrações disponíveis
// @Tags         v1 - Integração
// @Security     BearerAuth
// @Produce      json
// @Success      200  {array}  dto.IntegracaoResponse
// @Failure      500  {object}  rest_err.RestErr
// @Router       /admin/v1/integration [get]
func (ic *IntegrationController) GetAllIntegrations(ctx *gin.Context) {
	integracoes, err := ic.service.GetAllIntegrations()
	if err != nil {
		restErr := rest_err.NewInternalServerError("Erro ao buscar integrações", []rest_err.Causes{
			rest_err.NewCause("service", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromIntegracaoModelList(integracoes))
}

// GetAllMarcas godoc
// @Summary      Listar marcas de integração
// @Description  Retorna todas as marcas de integração
// @Tags         v1 - Integração
// @Security     BearerAuth
// @Produce      json
// @Success      200  {array}  dto.MarcaResponse
// @Failure      500  {object}  rest_err.RestErr
// @Router       /admin/v1/integration/marcas [get]
func (ic *IntegrationController) GetAllMarcas(ctx *gin.Context) {
	marcas, err := ic.service.GetAllMarcas()
	if err != nil {
		restErr := rest_err.NewInternalServerError("Erro ao buscar marcas", []rest_err.Causes{
			rest_err.NewCause("service", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromMarcaModelList(marcas))
}
