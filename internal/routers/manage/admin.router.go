package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (ar *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	//public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
		adminRouterPublic.POST("/otp")

	}
	// private router
	adminRouterPrivate := Router.Group("/admin/user")
	/* User liminiter Ahthen and Permission*/
	{
		adminRouterPrivate.GET("/active_user")
	}

}
