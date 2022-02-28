package middleware

import (
	"net/http"
	helper "rail/helpers/user"

	"github.com/gin-gonic/gin"
)

func AuthenticateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "token not found"})
			return
		}
		claims, err := helper.VerifyToken(clientToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Set("email", claims.Email)
		c.Set("firstName", claims.FirstName)
		c.Set("lastName", claims.LastName)
		c.Set("userId", claims.User_id)

		c.Next()
	}
}
