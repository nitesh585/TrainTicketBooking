package main

import (
	"invoiceService/docs"
	"invoiceService/gokafka"
	"invoiceService/logger"
	"invoiceService/middleware"
	"invoiceService/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var log logrus.Logger = *logger.GetLogger()

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Logging(), gin.Recovery())

	routes.InvoiceRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.WithFields(logrus.Fields{"err": err.Error()}).Error("Failed to load .env file")
	}
	docs.SwaggerInfo.Title = "Swagger Train-Ticket Booking System API"

	PORT := os.Getenv("PORT")
	log.WithFields(logrus.Fields{"Port": PORT}).Info("server listening on this port")

	router := setupRouter()

	go gokafka.InvoiceConsumer()

	router.Run(":" + PORT)
}
