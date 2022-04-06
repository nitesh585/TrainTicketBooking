package routes

import (
	"invoiceService/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(gin *gin.Engine) {
	i := gin.Group("/invoice")
	{
		i.GET("/getInvoice", controllers.GetInvoice())
	}
}
