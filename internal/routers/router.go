package routers

import (

	//c = controller
	c "GO-CRM-API-SHOPDEV/internal/controller"
	"GO-CRM-API-SHOPDEV/internal/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewServer() *gin.Engine {

	r := gin.Default()
	// Add middleware ~~
	r.Use(middlewares.AuthenMiddleware())

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", Pong)
		v1.GET("/user", c.NewUserController().GetUserById)
	}
	v2 := r.Group("/v2/2024")
	{
		v2.GET("/ping", Pong)
	}

	return r
}

func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "TanTD")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": " Xin chao: " + name,
		"uid":     uid,
	})
}
