package controller

import (
	"GO-CRM-API-SHOPDEV/internal/service"
	"GO-CRM-API-SHOPDEV/pkg/response"
	"github.com/gin-gonic/gin"
)

//type UserController struct {
//	userService *service.UserService
//}
//
//func NewUserController() *UserController {
//	return &UserController{
//
//		userService: service.NewUserService(),
//	}
//}
//
//// uc  user contr0ller
//func (uc *UserController) GetUserById(c *gin.Context) {
//
//	fmt.Println("-------> Handler Controller")
//	//	response.ErrorResponse(c, 203, "")
//	response.SuccessResponse(c, 201, uc.userService.GetInfoUseService())
//}

type UserController struct {
	userService service.IUserservice
}

func NewUserController(userService service.IUserservice) *UserController {
	return &UserController{
		userService: userService,
	}
}
func (uc *UserController) Register(c *gin.Context) {

	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, nil)

	//fmt.Println("-------> Handler Controller")
	////	response.ErrorResponse(c, 203, "")
	//response.SuccessResponse(c, 201, uc.userService.GetInfoUseService())
}
