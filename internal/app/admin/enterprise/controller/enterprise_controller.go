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

// Create godoc
// @Summary      Criar empresa
// @Description  Cria uma nova empresa com CNPJ e nome
// @Tags         v1 - Empresa
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateEnterpriseDTO  true  "Dados da empresa"
// @Success      201      {object}  dto.EnterpriseResponseDTO
// @Failure      400      {object}  rest_err.RestErr
// @Failure      500      {object}  rest_err.RestErr
// @Router       /admin/v1/enterprise [post]
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

// ReadByCNPJ godoc
// @Summary      Buscar empresa por CNPJ
// @Description  Retorna os dados de uma empresa com base no CNPJ fornecido.
// @Tags         v1 - Empresa
// @Accept       json
// @Produce      json
// @Param        cnpj  path      string  true  "CNPJ da empresa"
// @Success      200   {object}  dto.EnterpriseResponseDTO
// @Failure      400   {object}  rest_err.RestErr
// @Failure      404   {object}  rest_err.RestErr
// @Router       /admin/v1/enterprise/cnpj/{cnpj} [get]
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

// ReadByNome godoc
// @Summary      Buscar empresa por nome
// @Description  Retorna os dados de uma empresa com base no nome fornecido.
// @Tags         v1 - Empresa
// @Accept       json
// @Produce      json
// @Param        nome  path      string  true  "Nome da empresa"
// @Success      200   {object}  dto.EnterpriseResponseDTO
// @Failure      400   {object}  rest_err.RestErr
// @Failure      404   {object}  rest_err.RestErr
// @Router       /admin/v1/enterprise/nome/{nome} [get]
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

// ReadByID godoc
// @Summary      Buscar empresa por ID
// @Description  Retorna os dados de uma empresa com base no ID.
// @Tags         v1 - Empresa
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "ID da empresa"
// @Success      200   {object}  dto.EnterpriseResponseDTO
// @Failure      400   {object}  rest_err.RestErr
// @Failure      404   {object}  rest_err.RestErr
// @Router       /admin/v1/enterprise/id/{id} [get]
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

// ReadAll godoc
// @Summary      Listar todas as empresas
// @Description  Lista todas as empresas com suporte a paginação
// @Tags         v1 - Empresa
// @Accept       json
// @Produce      json
// @Param        page  query     int  true  "Número da página"
// @Success      200   {array}   dto.EnterpriseResponseDTO
// @Failure      400   {object}  rest_err.RestErr
// @Failure      404   {object}  rest_err.RestErr
// @Router       /admin/v1/enterprise [get]
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

// UpdateByCNPJ godoc
// @Summary      Atualizar empresa por CNPJ
// @Description  Atualiza os dados de uma empresa a partir do CNPJ fornecido na URL
// @Tags         v1 - Empresa
// @Accept       json
// @Produce      json
// @Param        cnpj    path      string                                 true  "CNPJ da empresa atual"
// @Param        request body      dto.UpdateEnterpriseByCNPJDTO         true  "Dados de atualização"
// @Success      200     {object}  dto.EnterpriseUpdatedResponseDTO
// @Failure      400     {object}  rest_err.RestErr
// @Failure      500     {object}  rest_err.RestErr
// @Router       /admin/v1/enterprise/cnpj/{cnpj} [put]
// Atualiza uma empresa com base no CNPJ original (URI) e dados novos (JSON)
func (ec *EnterpriseController) UpdateByCNPJ(ctx *gin.Context) {
	var req dto.UpdateEnterpriseByCNPJDTO

	// Bind do CNPJ da URI
	if err := ctx.ShouldBindUri(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("CNPJ ausente ou inválido na URL", []rest_err.Causes{
			rest_err.NewCause("uri", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Bind do corpo JSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Dados do corpo ausentes ou inválidos", []rest_err.Causes{
			rest_err.NewCause("body", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Validação personalizada
	if err := binding.ValidateUpdateEnterpriseDTO(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro de validação nos dados de atualização", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Monta o model com nome e cnpj novos (ou vazios)
	model := dto.UpdateEnterpriseByCNPJDTO{
		Cnpj: req.NewCNPJ,
		Nome: req.Nome,
	}.UpdateToModel()

	// Chama o service
	updated, err := ec.service.UpdateByCNPJ(req.Cnpj, req.NewCNPJ, &model)
	if err != nil {
		restErr := rest_err.NewInternalServerError("Erro ao atualizar empresa", []rest_err.Causes{
			rest_err.NewCause("repository", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Retorno mais limpo (nome, cnpj e update_at)
	ctx.JSON(http.StatusOK, dto.FromModelUpdateResponse(*updated))
}

// DeleteByCNPJ godoc
// @Summary      Deletar empresa por CNPJ
// @Description  Remove uma empresa com base no CNPJ fornecido.
// @Tags         v1 - Empresa
// @Accept       json
// @Produce      json
// @Param        request  path  dto.DeleteEnterpriseByCNPJDTO  true  "CNPJ para exclusão"
// @Success      200   {object}  dto.EnterpriseDeletedResponseDTO
// @Failure      400   {object}  rest_err.RestErr
// @Failure      404   {object}  rest_err.RestErr
// @Router       /admin/v1/enterprise/cnpj/{cnpj} [delete]
func (ec *EnterpriseController) DeleteByCNPJ(ctx *gin.Context) {
	var req dto.DeleteEnterpriseByCNPJDTO

	// Bind do CNPJ da URI
	if err := ctx.ShouldBindUri(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("CNPJ ausente ou inválido na URL", []rest_err.Causes{
			rest_err.NewCause("uri", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Validação personalizada
	if err := binding.ValidateDeleteEnterpriseByCNPJDTO(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro de validação no CNPJ", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Executa exclusão no service
	deletedCNPJ, err := ec.service.DeleteByCNPJ(req.Cnpj)
	if err != nil {
		restErr := rest_err.NewInternalServerError("Erro ao deletar empresa", []rest_err.Causes{
			rest_err.NewCause("repository", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Retorno de confirmação
	ctx.JSON(http.StatusOK, dto.EnterpriseDeletedResponseDTO{Cnpj: deletedCNPJ})
}
