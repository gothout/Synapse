package controller

import (
	binding "Synapse/internal/app/admin/binding/integration"
	dto "Synapse/internal/app/admin/integration/dto"
	integrationService "Synapse/internal/app/admin/integration/service"
	"Synapse/internal/configuration/rest_err"
	"net/http"
	"strings"

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

// DeleteIntegracaoFromEnterprise godoc
// @Summary      Remover vínculo entre empresa e integração
// @Description  Remove o vínculo entre uma empresa e uma integração existente
// @Tags         v1 - Integração
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body dto.DeleteIntegracaoEnterpriseRequest true "Dados para remover vínculo"
// @Success      204  {string} string "sem conteúdo"
// @Failure      400  {object} rest_err.RestErr
// @Failure      500  {object} rest_err.RestErr
// @Router       /admin/v1/integration/enterprise [delete]
func (ic *IntegrationController) DeleteIntegracaoFromEnterprise(ctx *gin.Context) {
	var req dto.DeleteIntegracaoEnterpriseRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro ao processar JSON de remoção", []rest_err.Causes{
			rest_err.NewCause("json", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if err := binding.ValidateDeleteIntegracaoEnterprise(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro de validação dos dados", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if err := ic.service.DeleteIntegracaoFromEnterprise(req.EnterpriseID, req.IntegracaoID); err != nil {
		restErr := rest_err.NewInternalServerError("Erro ao remover vínculo", []rest_err.Causes{
			rest_err.NewCause("service", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// CreateIntegracaoUser godoc
// @Summary      Vincular integração a um usuário
// @Description  Cria o vínculo entre um usuário e uma integração
// @Tags         v1 - Integração
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body dto.CreateIntegracaoUserRequest true "Dados para vincular usuário à integração"
// @Success      201  {string} string "vínculo criado com sucesso"
// @Failure      400  {object} rest_err.RestErr
// @Failure      403  {object} rest_err.RestErr
// @Failure      404  {object} rest_err.RestErr
// @Failure      500  {object} rest_err.RestErr
// @Router       /admin/v1/integration/user [post]
func (ic *IntegrationController) CreateIntegracaoUser(ctx *gin.Context) {
	var req dto.CreateIntegracaoUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro no JSON recebido", []rest_err.Causes{
			rest_err.NewCause("json", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if err := binding.ValidateCreateIntegracaoUser(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro de validação dos campos", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	err := ic.service.CreateIntegracaoUser(req.ToModel())
	if err != nil {
		// Tratamento específico baseado na mensagem do service
		switch {
		case strings.Contains(err.Error(), "usuário com ID"):
			ctx.JSON(http.StatusNotFound, rest_err.NewNotFoundError("Usuário não encontrado"))
			return
		case strings.Contains(err.Error(), "integração com ID"):
			ctx.JSON(http.StatusNotFound, rest_err.NewNotFoundError("Integração não encontrada"))
			return
		case strings.Contains(err.Error(), "violates unique constraint"): // opcional
			ctx.JSON(http.StatusForbidden, rest_err.NewForbiddenError("Este vínculo já existe"))
			return
		default:
			ctx.JSON(http.StatusInternalServerError, rest_err.NewInternalServerError("Erro ao criar vínculo usuário ↔ integração", []rest_err.Causes{
				rest_err.NewCause("service", err.Error()),
			}))
			return
		}
	}

	ctx.Status(http.StatusCreated)
}

// CreateTokenIntegracao godoc
// @Summary      Gerar token de Integração
// @Description  Autentica usuário admin e gera um token exclusivo para a integração
// @Tags         v1 - Integração
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body dto.CreateTokenIntegracaoRequest true "Dados de login e integração"
// @Success      200  {object} map[string]string "token de integração"
// @Failure      400  {object} rest_err.RestErr
// @Failure      403  {object} rest_err.RestErr
// @Failure      404  {object} rest_err.RestErr
// @Failure      500  {object} rest_err.RestErr
// @Router       /admin/v1/integration/token [post]
func (ic *IntegrationController) CreateTokenIntegracao(ctx *gin.Context) {
	var req dto.CreateTokenIntegracaoRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, rest_err.NewBadRequestValidationError("Erro no JSON recebido", []rest_err.Causes{
			rest_err.NewCause("json", err.Error()),
		}))
		return
	}

	if err := binding.ValidateCreateTokenIntegracao(req); err != nil {
		ctx.JSON(http.StatusBadRequest, rest_err.NewBadRequestValidationError("Erro de validação", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		}))
		return
	}

	token, err := ic.service.CreateTokenIntegracao(req.Email, req.Senha, req.IntegracaoID)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "credenciais"):
			ctx.JSON(http.StatusForbidden, rest_err.NewForbiddenError("Credenciais inválidas"))
			return
		case strings.Contains(err.Error(), "não possui permissão"):
			ctx.JSON(http.StatusForbidden, rest_err.NewForbiddenError("Usuário não vinculado à integração"))
			return
		case strings.Contains(err.Error(), "não encontrada"):
			ctx.JSON(http.StatusNotFound, rest_err.NewNotFoundError(err.Error()))
			return
		default:
			ctx.JSON(http.StatusInternalServerError, rest_err.NewInternalServerError("Erro ao gerar token", []rest_err.Causes{
				rest_err.NewCause("service", err.Error()),
			}))
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// GetIntegracoesByUserID godoc
// @Summary      Listar permissões de integração do usuário
// @Description  Retorna todas as integrações vinculadas a um usuário
// @Tags         v1 - Integração
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        user_id path int true "ID do usuário"
// @Success      200  {array}  dto.IntegracaoUsuarioResponse
// @Failure      400  {object} rest_err.RestErr
// @Failure      404  {object} rest_err.RestErr
// @Failure      500  {object} rest_err.RestErr
// @Router       /admin/v1/integration/user/{user_id} [get]
func (ic *IntegrationController) GetIntegracoesByUserID(ctx *gin.Context) {
	var req dto.GetIntegracoesByUserIDURI

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, rest_err.NewBadRequestValidationError("Erro ao processar URI", []rest_err.Causes{
			rest_err.NewCause("uri", err.Error()),
		}))
		return
	}

	if err := binding.ValidateGetIntegracoesByUserIDURI(req); err != nil {
		ctx.JSON(http.StatusBadRequest, rest_err.NewBadRequestValidationError("Erro de validação", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		}))
		return
	}

	integs, err := ic.service.GetIntegracoesByUserID(req.ToID())
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "usuário"):
			ctx.JSON(http.StatusNotFound, rest_err.NewNotFoundError("Usuário não encontrado"))
		default:
			ctx.JSON(http.StatusInternalServerError, rest_err.NewInternalServerError("Erro ao buscar permissões de integração", []rest_err.Causes{
				rest_err.NewCause("service", err.Error()),
			}))
		}
		return
	}

	ctx.JSON(http.StatusOK, dto.FromIntegracaoUsuarioModelList(integs))
}

// DeleteIntegracaoUsuario godoc
// @Summary      Remover vínculo entre usuário e integração
// @Description  Remove o vínculo entre um usuário e uma integração existente
// @Tags         v1 - Integração
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body dto.RemoveIntegracaoUsuarioRequest true "Dados para remover vínculo"
// @Success      204  {string} string "sem conteúdo"
// @Failure      400  {object} rest_err.RestErr
// @Failure      404  {object} rest_err.RestErr
// @Failure      500  {object} rest_err.RestErr
// @Router       /admin/v1/integration/user [delete]
func (ic *IntegrationController) DeleteIntegracaoUsuario(ctx *gin.Context) {
	var req dto.RemoveIntegracaoUsuarioRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, rest_err.NewBadRequestValidationError("Erro ao processar JSON", []rest_err.Causes{
			rest_err.NewCause("json", err.Error()),
		}))
		return
	}

	if err := binding.ValidateRemoveIntegracaoUsuario(req); err != nil {
		ctx.JSON(http.StatusBadRequest, rest_err.NewBadRequestValidationError("Erro de validação dos dados", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		}))
		return
	}

	if err := ic.service.RemoveIntegrationFromUser(ctx, req.UserID, req.IntegrationID); err != nil {
		switch {
		case strings.Contains(err.Error(), "usuário"):
			ctx.JSON(http.StatusNotFound, rest_err.NewNotFoundError("Usuário não encontrado"))
			return
		case strings.Contains(err.Error(), "integração"):
			ctx.JSON(http.StatusNotFound, rest_err.NewNotFoundError("Integração não encontrada"))
			return
		case strings.Contains(err.Error(), "nenhum vínculo encontrado"):
			ctx.JSON(http.StatusNotFound, rest_err.NewNotFoundError("Vínculo não encontrado"))
			return
		default:
			ctx.JSON(http.StatusInternalServerError, rest_err.NewInternalServerError("Erro ao remover vínculo usuário ↔ integração", []rest_err.Causes{
				rest_err.NewCause("service", err.Error()),
			}))
			return
		}
	}

	ctx.Status(http.StatusNoContent)
}
