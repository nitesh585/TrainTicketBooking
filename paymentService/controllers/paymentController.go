package controllers

import (
	"net/http"
	"os"
	paymentHelper "paymentService/helpers"
	"paymentService/logger"
	models "paymentService/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	razorpay "github.com/razorpay/razorpay-go"
	"github.com/sirupsen/logrus"
)

var log logrus.Logger = *logger.GetLogger()

// ShowAccount godoc
// @Summary      Pay for train ticket
// @Description  Get razorpay payment link along with user details for ticket
// @Tags         payment
// @Accept       json
// @Produce      json
// @Param        Date  			body 	string  true  "date of Booking"
// @Param        Amount 		body	int   	true  "amount of tickets"
// @Param        Currency 		body	string  true  "type of currency in which transaction will happend"
// @Param        TrainID 		body	string  true  "train number"
// @Param        ReferenceID 	body	string  false "reference id is a unique id for payment gatway"
// @Param        ClassBooking 	body	string  true  "to which class the tickets belongs, e.g, SL, AC-1, AC-2, etc"
// @Param        Name 			body	string  true  "user name"
// @Param        Email 			body	string  true  "user email id"
// @Param        User_id 		body	string  true  "user unique id"
// @Success      200  {object}  models.PaymentResponse
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /payment/pay [post]
func PayForTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Load(".env")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"err": err.Error()}).Error("Failed to load .env file")
			return
		}

		key_id := os.Getenv("RAZORPAY_KEY_ID")
		key_secret := os.Getenv("RAZORPAY_KEY_SECRET")

		var payTicketDetails models.PayTicket
		if err := c.BindJSON(&payTicketDetails); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"err": err.Error()}).Error("fail to bind body to gin json")
			return
		}

		referenceID, err := paymentHelper.RandRefrenceID(8)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"err": err.Error()}).
				Error("fail to create new refrence id")
			return
		}

		log.Debug("referenceId created successfully")

		client := razorpay.NewClient(key_id, key_secret)
		log.Debug("razorpay client is created")

		data := gin.H{
			"amount":       payTicketDetails.Amount,
			"currency":     payTicketDetails.Currency,
			"reference_id": referenceID,
			"customer": models.PayCustomerDetails{
				Name:  payTicketDetails.Name,
				Email: payTicketDetails.Email,
			},
			"notes": models.PayNotes{
				UserID:        payTicketDetails.User_id,
				DateOfBooking: payTicketDetails.Date,
				ClassBooked:   payTicketDetails.ClassBooking,
				TrainId:       payTicketDetails.TrainID,
			},
		}
		body, err := client.PaymentLink.Create(data, nil)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err, "body": body})
			log.WithFields(logrus.Fields{"err": err.Error()}).Error("fail to create payment link")
			return
		}
		c.JSON(http.StatusOK, gin.H{"link": body})
		log.Info("payment link successfully sent")
	}
}
