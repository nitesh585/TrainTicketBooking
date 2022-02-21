package routes

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.Engine) {
	u := r.Group("/user")
	{
		u.GET("register", func(c *gin.Context) {

		})
	}

	t := r.Group("/train")
	{
		t.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Train",
			})
		})
	}

}
