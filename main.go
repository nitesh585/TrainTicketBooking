package main

import (
	"rail/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// routes.SetupRouter(r)

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
