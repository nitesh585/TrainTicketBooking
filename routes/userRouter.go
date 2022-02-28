package routes

import (
	controller "rail/controllers/user"
	"rail/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(gin *gin.Engine) {
	gin.Use(middleware.AuthenticateToken())

	u := gin.Group("/user")
	{
		u.GET("", controller.GetUsers)
		u.GET(":user_id", controller.GetUser)
	}
}
