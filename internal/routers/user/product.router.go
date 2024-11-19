package user

import "github.com/gin-gonic/gin"

type ProductRouter struct {
}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	//public router
	productRoutersPublic := Router.Group("/product")
	{
		productRoutersPublic.GET("/search")
		productRoutersPublic.GET("/detail")
	}

}
