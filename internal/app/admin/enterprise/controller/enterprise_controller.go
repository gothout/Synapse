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

// Busca empresa por CNPJ
func (ec *EnterpriseController) ReadByCNPJ(ctx *gin.Context) {
	var req dto.ReadEnterpriseCNPJDTO

	// Bind automático do parâmetro de rota para DTO
	if err := ctx.ShouldBindUri(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("CNPJ ausente ou inválido na URL", []rest_err.Causes{
			rest_err.NewCause("uri", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Validação personalizada
	if err := binding.ValidateReadEnterpriseCNPJDTO(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro de validação no CNPJ", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	enterprise, err := ec.service.ReadByCNPJ(req.Cnpj)
	if err != nil {
		restErr := rest_err.NewNotFoundError("Empresa não encontrada")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromModel(*enterprise))
}

// Busca empresa por NOME
func (ec *EnterpriseController) ReadByNome(ctx *gin.Context) {
	var req dto.ReadEnterpriseNOMEDTO

	// Bind automático do parâmetro de rota para DTO
	if err := ctx.ShouldBindUri(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("NOME ausente ou inválido na URL", []rest_err.Causes{
			rest_err.NewCause("uri", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	enterprise, err := ec.service.ReadByNome(req.Nome)
	if err != nil {
		restErr := rest_err.NewNotFoundError("Empresa não encontrada")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromModel(*enterprise))
}

// Busca empresa por ID
func (ec *EnterpriseController) ReadByID(ctx *gin.Context) {
	var req dto.ReadEnterpriseIDDTO

	// Bind automático do parâmetro de rota para DTO
	if err := ctx.ShouldBindUri(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("ID ausente ou inválido na URL", []rest_err.Causes{
			rest_err.NewCause("uri", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	enterprise, err := ec.service.ReadByID(req.ID)
	if err != nil {
		restErr := rest_err.NewNotFoundError("Empresa não encontrada")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromModel(*enterprise))
}

// Busca empresas por Page
func (ec *EnterpriseController) ReadAll(ctx *gin.Context) {
	var req dto.ReadEnterpriseAllDTO

	// Bind automático do query param para DTO
	if err := ctx.ShouldBindQuery(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Page deve ser um valor numérico", []rest_err.Causes{
			rest_err.NewCause("query", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	page, err := binding.ValidateReadEnterpriseAllDTO(req)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro na validação da página", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	empresas, err := ec.service.ReadAll(page)
	if err != nil {
		restErr := rest_err.NewNotFoundError("Nenhuma empresa encontrada")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromModelList(empresas))
}
