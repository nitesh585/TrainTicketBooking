package main

import (
	"log"
	"os"
	"rail/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file", err.Error())
	}

	PORT := os.Getenv("PORT")

	router := gin.Default()
	// routes.SetupRouter(r)

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + PORT) // listen and serve on 0.0.0.0:8080
}
