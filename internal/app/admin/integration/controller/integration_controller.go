package controller

import (
	binding "Synapse/internal/app/admin/binding/integration"
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

// GetDetalhadasByMarcaID godoc
// @Summary      Listar integrações detalhadas por marca
// @Description  Retorna todas as integrações vinculadas a uma marca com nome da empresa
// @Tags         v1 - Integração
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        marca_id query int true "ID da marca"
// @Success      200  {array}  dto.IntegracaoDetalhadaResponse
// @Failure      400  {object} rest_err.RestErr
// @Failure      500  {object} rest_err.RestErr
// @Router       /admin/v1/integration/marca/detalhada [get]
func (ic *IntegrationController) GetDetalhadasByMarcaID(ctx *gin.Context) {
	var req dto.GetIntegracoesByMarcaIDRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro ao processar query", []rest_err.Causes{
			rest_err.NewCause("query", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if err := binding.ValidateGetIntegracoesByMarcaID(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("marca_id inválido ou ausente", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	integracoes, err := ic.service.GetIntegrationsByMarcaID(req.MarcaID)
	if err != nil {
		restErr := rest_err.NewInternalServerError("Erro ao buscar integrações detalhadas", []rest_err.Causes{
			rest_err.NewCause("service", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromIntegracaoDetalhadaModelList(integracoes))
}

// CreateIntegracaoEnterprise godoc
// @Summary      Vincular integração a uma empresa
// @Description  Cria o vínculo entre uma empresa e uma integração
// @Tags         v1 - Integração
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body dto.CreateIntegracaoEnterpriseRequest true "Dados para vincular empresa à integração"
// @Success      201  {string} string "criado com sucesso"
// @Failure      400  {object} rest_err.RestErr
// @Failure      500  {object} rest_err.RestErr
// @Router       /admin/v1/integration/enterprise [post]
func (ic *IntegrationController) CreateIntegracaoEnterprise(ctx *gin.Context) {
	var req dto.CreateIntegracaoEnterpriseRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro ao processar JSON", []rest_err.Causes{
			rest_err.NewCause("json", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if err := binding.ValidateCreateIntegracaoEnterprise(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Dados inválidos", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if err := ic.service.CreateIntegracaoEnterprise(req.ToModelIntegracaoEnterprise()); err != nil {
		restErr := rest_err.NewInternalServerError("Erro ao vincular integração à empresa", []rest_err.Causes{
			rest_err.NewCause("service", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.Status(http.StatusCreated)
}

// GetByEnterpriseID godoc
// @Summary      Listar integrações liberadas para uma empresa
// @Description  Retorna todas as integrações vinculadas a uma empresa específica
// @Tags         v1 - Integração
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        enterprise_id path int true "ID da empresa"
// @Success      200  {array}  dto.IntegracaoEmpresaDetalhadaResponse
// @Failure      400  {object} rest_err.RestErr
// @Failure      500  {object} rest_err.RestErr
// @Router       /admin/v1/integration/enterprise/{enterprise_id} [get]
func (ic *IntegrationController) GetByEnterpriseID(ctx *gin.Context) {
	var req dto.GetIntegracoesByEnterpriseIDURI

	// Faz o bind da URI -> struct
	if err := ctx.ShouldBindUri(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro ao extrair enterprise_id da rota", []rest_err.Causes{
			rest_err.NewCause("uri", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Valida valor da struct
	if err := binding.ValidateGetIntegracoesByEnterpriseURI(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro de validação da rota", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Chama o service
	integracoes, err := ic.service.GetIntegracoesByEnterpriseID(req.ToID())
	if err != nil {
		restErr := rest_err.NewInternalServerError("Erro ao buscar integrações da empresa", []rest_err.Causes{
			rest_err.NewCause("service", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Resposta final
	ctx.JSON(http.StatusOK, dto.FromIntegracaoEmpresaDetalhadaModelList(integracoes))
}
