package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public router
	// this is non-dependency

	UserRouterPublic := Router.Group("/user")
	{
		UserRouterPublic.POST("/register")
		//	UserRouterPublic.POST("/opt")
	}
	// private router
	UserRouterPrivate := Router.Group("/user")
	/* User liminiter Ahthen and Permission*/
	{
		UserRouterPrivate.GET("/get_info")
		UserRouterPrivate.POST("/")
	}
	// private router

}
