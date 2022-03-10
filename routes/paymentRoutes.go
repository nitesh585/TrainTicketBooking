package routes

import (
	controller "rail/controllers/payment"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(gin *gin.Engine) {
	p := gin.Group("/payment")
	{
		p.GET("/pay", controller.PayForTicket())
	}
}
