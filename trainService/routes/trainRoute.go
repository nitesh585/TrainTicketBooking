package routes

import (
	"trainService/controllers"
	"trainService/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log logrus.Logger = *logger.GetLogger()

func TrainRouter(gin *gin.Engine) {
	t := gin.Group("/train")
	{
		t.POST("/checkAvailability", controllers.CheckAvailability())
		t.GET("/searchRoute", controllers.SearchRoute())
		t.GET("/trainDetails", controllers.TrainDetails())
	}
	log.Trace("train routes are added. ")
}
