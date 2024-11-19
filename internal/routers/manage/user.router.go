package manage

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public router
	// private router
	UserRouterPrivate := Router.Group("/admin/user")
	/* User liminiter Ahthen and Permission*/
	{
		UserRouterPrivate.POST("/active_user")
	}
	// private router

}
