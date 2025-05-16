package user

import (
	binding "Synapse/internal/app/admin/binding/user"
	dto "Synapse/internal/app/admin/user/dto"
	userService "Synapse/internal/app/admin/user/service"
	"Synapse/internal/configuration/rest_err"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service userService.Service
}

func NewUserController(s userService.Service) *UserController {
	return &UserController{service: s}
}

// Create godoc
// @Summary      Criar usuário
// @Description  Cria um novo usuário no sistema
// @Tags         v1 - Usuário
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AdminUserCreateDTO  true  "Dados do usuário"
// @Success      201      {object}  dto.UserResponseDTO
// @Failure      400      {object}  rest_err.RestErr
// @Failure      500      {object}  rest_err.RestErr
// @Router       /admin/v1/user [post]
func (uc *UserController) Create(ctx *gin.Context) {
	var req dto.AdminUserCreateDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Campos obrigatórios ausentes ou inválidos", []rest_err.Causes{
			rest_err.NewCause("body", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if err := binding.ValidateAdminUserCreateDTO(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro de validação nos dados do usuário", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	model := req.ToModel()
	created, err := uc.service.Create(&model)
	if err != nil {
		restErr := rest_err.NewInternalServerError("Erro ao criar usuário", []rest_err.Causes{
			rest_err.NewCause("Erro ao criar usuário", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusCreated, dto.FromModel(*created))
}

// ReadAll godoc
// @Summary      Listar usuários por empresa
// @Description  Lista todos os usuários de uma empresa com paginação
// @Tags         v1 - Usuário
// @Accept       json
// @Produce      json
// @Param        enterprise_id  path   int     true  "ID da empresa"
// @Param        page           query  string  true  "Número da página"
// @Success      200  {array}   dto.UserResponseDTO
// @Failure      400  {object}  rest_err.RestErr
// @Failure      404  {object}  rest_err.RestErr
// @Router       /admin/v1/user/{enterprise_id} [get]
func (uc *UserController) ReadAll(ctx *gin.Context) {
	var uri dto.AdminUserReadAllURI
	var query dto.AdminUserReadAllQuery

	// Bind da URI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Parâmetro enterprise_id inválido", []rest_err.Causes{
			rest_err.NewCause("uri", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Bind da query
	if err := ctx.ShouldBindQuery(&query); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Parâmetro page inválido", []rest_err.Causes{
			rest_err.NewCause("query", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Validação manual
	if err := binding.ValidateAdminUserReadAllDTO(uri.EnterpriseID, query.Page); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro de validação nos parâmetros", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Conversão do page para inteiro
	pageInt, err := strconv.Atoi(query.Page)
	if err != nil || pageInt <= 0 {
		restErr := rest_err.NewBadRequestValidationError("Parâmetro page deve ser um número válido maior que zero", []rest_err.Causes{
			rest_err.NewCause("page", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Chamada do service
	users, err := uc.service.ReadAllUser(uri.EnterpriseID, int64(pageInt))
	if err != nil {
		restErr := rest_err.NewNotFoundError("Nenhum usuário encontrado")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromModelList(*users))
}

// ReadByEmail godoc
// @Summary      Buscar usuário por e-mail
// @Description  Retorna os dados de um usuário com base no e-mail fornecido
// @Tags         v1 - Usuário
// @Accept       json
// @Produce      json
// @Param        email  path      string  true  "E-mail do usuário"
// @Success      200    {object}  dto.UserResponseDTO
// @Failure      400    {object}  rest_err.RestErr
// @Failure      404    {object}  rest_err.RestErr
// @Router       /admin/v1/user/email/{email} [get]
func (uc *UserController) ReadByEmail(ctx *gin.Context) {
	var req dto.AdminUserReadByEmailDTO

	if err := ctx.ShouldBindUri(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("E-mail inválido na rota", []rest_err.Causes{
			rest_err.NewCause("uri", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if err := binding.ValidateAdminUserReadByEmailDTO(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro de validação no e-mail", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	user, err := uc.service.ReadByEmail(req.Email)
	if err != nil {
		restErr := rest_err.NewNotFoundError("Usuário não encontrado")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromModel(*user))
}
