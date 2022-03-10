package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	razorpay "github.com/razorpay/razorpay-go"
)

func PayForTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Load(".env")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		key_id := os.Getenv("RAZORPAY_KEY_ID")
		key_secret := os.Getenv("RAZORPAY_KEY_SECRET")

		client := razorpay.NewClient(key_id, key_secret)
		data := gin.H{
			"amount":   50000,
			"currency": "INR",
			"receipt":  "some_receipt_id",
		}
		body, err := client.Order.Create(data, nil)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err, "body": body})
			return
		}

		c.JSON(http.StatusOK, gin.H{"link": body})
	}
}
