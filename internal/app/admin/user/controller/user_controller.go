package user

import (
	binding "Synapse/internal/app/admin/binding/user"
	dto "Synapse/internal/app/admin/user/dto"
	userService "Synapse/internal/app/admin/user/service"
	"Synapse/internal/configuration/rest_err"
	"fmt"
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

// Update godoc
// @Summary      Atualizar usuário
// @Description  Atualiza os dados de um usuário com base no ID
// @Tags         v1 - Usuário
// @Accept       json
// @Produce      json
// @Param        id       path      int                     true  "ID do usuário"
// @Param        request  body      dto.AdminUserUpdatedDTO  true  "Dados do usuário para atualização"
// @Success      200      {object}  dto.UserResponseDTO
// @Failure      400      {object}  rest_err.RestErr
// @Failure      404      {object}  rest_err.RestErr
// @Failure      500      {object}  rest_err.RestErr
// @Router       /admin/v1/user/{id} [put]
func (uc *UserController) Update(ctx *gin.Context) {
	// Extrai o ID da URL
	idParam := ctx.Param("id")
	userID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || userID <= 0 {
		restErr := rest_err.NewBadRequestValidationError("ID do usuário inválido", []rest_err.Causes{
			rest_err.NewCause("id", "ID deve ser um número inteiro válido"),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Bind do JSON
	var req dto.AdminUserUpdatedDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Campos obrigatórios ausentes ou inválidos", []rest_err.Causes{
			rest_err.NewCause("body", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Validação manual
	if err := binding.ValidateAdminUserUpdateDTO(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro de validação nos dados do usuário", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Converte DTO para model
	updatedUser := req.ToModelUpdated()

	// Chamada do service
	user, err := uc.service.UpdateUserByID(userID, &updatedUser)
	if err != nil {
		if err.Error() == fmt.Sprintf("usuário com ID %d não encontrado", userID) {
			restErr := rest_err.NewNotFoundError("Usuário não encontrado")
			ctx.JSON(restErr.Code, restErr)
			return
		}
		restErr := rest_err.NewInternalServerError("Erro ao atualizar usuário", []rest_err.Causes{
			rest_err.NewCause("update", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Retorno formatado
	ctx.JSON(http.StatusOK, dto.FromModelUpdated(*user))
}

// Delete godoc
// @Summary      Deletar usuário
// @Description  Remove um usuário com base no ID fornecido
// @Tags         v1 - Usuário
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do usuário"
// @Success      204  "Nenhum conteúdo"
// @Failure      400  {object}  rest_err.RestErr
// @Failure      404  {object}  rest_err.RestErr
// @Failure      500  {object}  rest_err.RestErr
// @Router       /admin/v1/user/{id} [delete]
func (uc *UserController) Delete(ctx *gin.Context) {
	var uri dto.AdminDeleteIDUserDTO

	// Bind do ID da URI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		restErr := rest_err.NewBadRequestValidationError("ID do usuário inválido na URI", []rest_err.Causes{
			rest_err.NewCause("uri", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Validação do ID
	if err := binding.ValidateAdminUserDeleteDTO(uri); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Erro de validação no ID do usuário", []rest_err.Causes{
			rest_err.NewCause("validação", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Execução da deleção
	if err := uc.service.DeleteUserByID(uri.UserID); err != nil {
		if err.Error() == fmt.Sprintf("nenhum usuário encontrado com ID %d", uri.UserID) {
			restErr := rest_err.NewNotFoundError("Usuário não encontrado")
			ctx.JSON(restErr.Code, restErr)
			return
		}
		restErr := rest_err.NewInternalServerError("Erro ao deletar usuário", []rest_err.Causes{
			rest_err.NewCause("delete", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
