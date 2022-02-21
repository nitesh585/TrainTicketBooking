package main

import (
	"rail/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
