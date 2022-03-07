package routes

import (
	controller "rail/controllers/train"

	"github.com/gin-gonic/gin"
)

func TrainRoute(gin *gin.Engine) {
	t := gin.Group("/train")
	{
		t.GET("/searchRoute", controller.SearchRoute())
	}
}
