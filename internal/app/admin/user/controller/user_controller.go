package user

import (
	binding "Synapse/internal/app/admin/binding/user"
	dto "Synapse/internal/app/admin/user/dto"
	userService "Synapse/internal/app/admin/user/service"
	"Synapse/internal/configuration/rest_err"
	"net/http"

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
