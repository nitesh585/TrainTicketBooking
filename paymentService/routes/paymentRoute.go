package routes

import (
	controller "paymentService/controllers"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(gin *gin.Engine) {
	p := gin.Group("/payment")
	{
		p.GET("/pay", controller.PayForTicket())
	}
}
