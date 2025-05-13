package rule

import (
	binding "Synapse/internal/app/admin/binding/rule"
	dto "Synapse/internal/app/admin/rule/dto"
	ruleService "Synapse/internal/app/admin/rule/service"
	"Synapse/internal/configuration/rest_err"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RuleController struct {
	service ruleService.Service
}

func NewRuleController(s ruleService.Service) *RuleController {
	return &RuleController{service: s}
}

// GetAll godoc
// @Summary      Listar regras
// @Description  Retorna todas as regras com suporte a paginação
// @Tags         v1 - Regra
// @Accept       json
// @Produce      json
// @Param        limit   query     int  true  "Quantidade de resultados"
// @Param        offset  query     int  true  "Deslocamento dos resultados"
// @Success      200  {object}  dto.AdminRuleListDTO
// @Failure      400  {object}  rest_err.RestErr
// @Router       /admin/v1/rules [get]
func (rc *RuleController) GetAll(ctx *gin.Context) {
	var req dto.GetRulesQueryDTO

	if err := ctx.ShouldBindQuery(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Parâmetros de paginação inválidos", []rest_err.Causes{
			rest_err.NewCause("query", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	limit, offset, err := binding.ValidateGetRulesQueryDTO(req)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro na validação da paginação", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	rules, err := rc.service.FindAllRules(limit, offset)
	if err != nil {
		restErr := rest_err.NewInternalServerError("Erro ao buscar regras", []rest_err.Causes{
			rest_err.NewCause("service", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	response := dto.AdminRuleListDTO{
		Total:  int64(len(rules)), // você pode ajustar futuramente com COUNT no banco
		Limit:  limit,
		Offset: offset,
		Data:   dto.FromModelList(rules),
	}

	ctx.JSON(http.StatusOK, response)
}

// GetByID godoc
// @Summary      Buscar regra por ID
// @Description  Retorna os dados de uma regra com base no ID
// @Tags         v1 - Regra
// @Accept       json
// @Produce      json
// @Param        id  path      int  true  "ID da regra"
// @Success      200  {object}  dto.AdminRuleDTO
// @Failure      400  {object}  rest_err.RestErr
// @Failure      404  {object}  rest_err.RestErr
// @Router       /admin/v1/rules/{id} [get]
func (rc *RuleController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("ID inválido na URL", []rest_err.Causes{
			rest_err.NewCause("param", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if err := binding.ValidateReadRuleByIDDTO(id); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro na validação do ID", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	rule, err := rc.service.FindRuleByID(id)
	if err != nil {
		restErr := rest_err.NewNotFoundError("Regra não encontrada")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromModel(*rule))
}

// GetPermissions godoc
// @Summary      Buscar permissões da regra
// @Description  Retorna a lista de permissões em formato namespace (ex: admin.enterprise.create)
// @Tags         v1 - Regra
// @Accept       json
// @Produce      json
// @Param        id  path      int  true  "ID da regra"
// @Success      200  {array}   dto.RulePermissionDTO
// @Failure      400  {object}  rest_err.RestErr
// @Failure      404  {object}  rest_err.RestErr
// @Router       /admin/v1/rules/{id}/permissions [get]
func (rc *RuleController) GetPermissions(ctx *gin.Context) {
	idParam := ctx.Param("id")
	ruleID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("ID inválido na URL", []rest_err.Causes{
			rest_err.NewCause("param", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if err := binding.ValidateReadPermissionsByRuleIDDTO(ruleID); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro na validação do ID da regra", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	namespaces, err := rc.service.FindPermissionsByRuleID(ruleID)
	if err != nil {
		restErr := rest_err.NewNotFoundError("Permissões não encontradas para a regra")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromNamespaceList(namespaces))
}
