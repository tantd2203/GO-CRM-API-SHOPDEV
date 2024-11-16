package controller

import (
	"GO-CRM-API-SHOPDEV/internal/service"
	"GO-CRM-API-SHOPDEV/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{

		userService: service.NewUserService(),
	}
}

// uc  user contr0ller
func (uc *UserController) GetUserById(c *gin.Context) {
	response.ErrorResponse(c, 203, "")
	//response.SuccessResponse(c, 201, uc.userService.GetInfoUseService())
}
