package controller

import (
	"GO-CRM-API-SHOPDEV/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{

		userService: service.NewUserSerivce(),
	}
}

// uc  user contr0ller
func (uc *UserController) GetUserById(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfoUseService(),
	})
}
