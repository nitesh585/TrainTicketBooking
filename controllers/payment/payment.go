package controller

import (
	"net/http"
	"os"
	paymentHelper "rail/helpers/payment"
	authHelper "rail/helpers/user"

	models "rail/models/payment"

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

		var payTicketDetails models.PayTicket
		if err := c.BindJSON(&payTicketDetails); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		clientToken := c.Request.Header.Get("token")
		claims, _ := authHelper.VerifyToken(clientToken)

		referenceID, err := paymentHelper.RandRefrenceID(8)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		client := razorpay.NewClient(key_id, key_secret)
		data := gin.H{
			"amount":       payTicketDetails.Amount,
			"currency":     payTicketDetails.Currency,
			"reference_id": referenceID,
			"customer": models.PayCustomerDetails{
				Name:  claims.FirstName + " " + claims.LastName,
				Email: claims.Email,
			},
			"notes": models.PayNotes{
				UserID:        claims.User_id,
				DateOfBooking: payTicketDetails.Date,
				ClassBooked:   payTicketDetails.ClassBooking,
				TrainId:       payTicketDetails.TrainID,
			},
		}
		body, err := client.PaymentLink.Create(data, nil)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err, "body": body})
			return
		}
		c.JSON(http.StatusOK, gin.H{"link": body})
	}
}
