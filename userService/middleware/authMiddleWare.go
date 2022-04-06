package middleware

import (
	"net/http"
	helper "userService/helpers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthenticateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "token not found"})
			log.Error("token not found")
			return
		}
		claims, err := helper.VerifyToken(clientToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).
				Error("verification of token failed")
			return
		}

		c.Set("email", claims.Email)
		c.Set("firstName", claims.FirstName)
		c.Set("lastName", claims.LastName)
		c.Set("userId", claims.User_id)

		log.Debug("user successfully authenticated")
		c.Next()
	}
}
