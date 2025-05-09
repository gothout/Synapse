package enterprise

import (
	binding "Synapse/internal/app/admin/binding/enterprise"
	dto "Synapse/internal/app/admin/enterprise/dto"
	enterpriseService "Synapse/internal/app/admin/enterprise/service"
	"Synapse/internal/configuration/rest_err"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EnterpriseController struct {
	service enterpriseService.Service
}

func NewEnterpriseController(s enterpriseService.Service) *EnterpriseController {
	return &EnterpriseController{service: s}
}

func (ec *EnterpriseController) Create(ctx *gin.Context) {
	var req dto.CreateEnterpriseDTO

	// Bind JSON + validação obrigatória (tags)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Campos obrigatórios ausentes ou inválidos", []rest_err.Causes{
			rest_err.NewCause("body", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Validação personalizada
	if err := binding.ValidateCreateEnterpriseDTO(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro de validação nos dados", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	model := req.ToModel()

	created, err := ec.service.Create(&model)
	if err != nil {
		// Você pode ajustar isso para diferenciar erros do service também
		restErr := rest_err.NewInternalServerError("Erro ao criar empresa", []rest_err.Causes{
			rest_err.NewCause("repository", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusCreated, dto.FromModel(*created))
}
