package main

import (
	"os"
	"paymentService/logger"
	"paymentService/middleware"
	"paymentService/routes"

	"paymentService/docs"

	"github.com/gin-gonic/gin" // swagger embed files
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var log logrus.Logger = *logger.GetLogger()

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Logging(), gin.Recovery())

	routes.PaymentRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

// @title           Swagger Train-Ticket Booking System API
// @version         1.0
// @description     Swagger Train-Ticket Booking System API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  swiggyb2013@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082

// @securityDefinitions.basic  BasicAuth
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.WithFields(logrus.Fields{"err": err.Error()}).Error("Failed to load .env file")
	}
	docs.SwaggerInfo.Title = "Swagger Train-Ticket Booking System API"

	PORT := os.Getenv("PORT")
	log.WithFields(logrus.Fields{"Port": PORT}).Info("server listening on this port")

	router := setupRouter()
	router.Run(":" + PORT)
}
