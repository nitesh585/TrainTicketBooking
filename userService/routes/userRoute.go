package routes

import (
	"userService/controllers"
	"userService/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log logrus.Logger = *logger.GetLogger()

func UserRouter(gin *gin.Engine) {
	u := gin.Group("/user")
	{
		u.POST("/signup", controllers.Signup())
		u.POST("/login", controllers.Login())
	}
	log.Trace("user routes are added. ")
}
