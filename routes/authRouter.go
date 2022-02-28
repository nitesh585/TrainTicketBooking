package routes

import (
	"github.com/gin-gonic/gin"

	controller "rail/controllers/user"
)

func AuthRoutes(gin *gin.Engine) {
	u := gin.Group("/user")
	{
		u.POST("/signup", controller.Signup())
		u.POST("/login", controller.Login())
	}
}
