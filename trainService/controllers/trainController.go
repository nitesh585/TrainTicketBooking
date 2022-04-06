package controllers

import (
	"context"
	"net/http"
	"os"
	"time"
	"trainService/database"
	helper "trainService/helpers"
	"trainService/logger"
	models "trainService/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var log logrus.Logger = *logger.GetLogger()

func getTrainCollection() *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.WithFields(logrus.Fields{"err": err.Error()}).Error("Failed to load .env file")
	}

	COLLECTION := os.Getenv("COLLECTION")
	trainCollection := database.OpenCollection(database.MongoClient, COLLECTION)

	return trainCollection
}

var trainCollection *mongo.Collection = getTrainCollection()

// ShowAccount godoc
// @Summary      Check Availability
// @Description  check whether train is available on particular date or not
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param        TrainID  body 	string  true  "unique train id"
// @Param        Date 	  body	string  true  "date of booking"
// @Success      200  {object} 	models.Train
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /train/checkAvailability [post]
func CheckAvailability() gin.HandlerFunc {
	return func(g *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
		defer cancel()

		var search struct {
			TrainID string
			Date    string
		}

		if err := g.BindJSON(&search); err != nil {
			g.JSON(http.StatusOK, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("failed to load .env file")
			return
		}

		log.Debug("gin json binding done.")

		var trainDetails models.Train
		err := trainCollection.FindOne(ctx, bson.M{"TrainID": search.TrainID}).Decode(&trainDetails)

		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": "no data found"})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("o data found")
			return
		}

		log.Debug("train ID found in DB.")
		log.Trace("train ID found in DB.")

		//-------------check for date availability
		layout := "01/02/06"
		parseDate, err := time.Parse(layout, search.Date)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("failed to parse date")
			return
		}

		log.Debug("date present in query is parsed.")
		log.Trace("date present in query is parsed.")

		if time.Now().After(parseDate) {
			g.JSON(http.StatusBadRequest, gin.H{"error": "invalid date"})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("invalid date")
			return
		}

		weekday := parseDate.Weekday().String()
		isAvailable := helper.IsTrainAvailableOnGivenWeekDay(weekday, trainDetails)
		if !isAvailable {
			g.JSON(http.StatusBadRequest, gin.H{"msg": "train not available"})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("train is not available")
			return
		}
		log.Debug("train is available on given date.")

		responseTrainDetails := helper.CalculatePriceOne(
			trainDetails.FromStationCode,
			trainDetails.ToStationCode,
			trainDetails,
		)

		g.JSON(http.StatusOK, gin.H{"trainDetails": responseTrainDetails})
	}
}

type SearchQuery struct {
	Source      string
	Destination string
	Date        string
}

// ShowAccount godoc
// @Summary      Search route from source to destination stations
// @Description  search route from source to destination stations on specific date
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param        Source  		body 	string  true  "source station code"
// @Param        Destination 	body	string  true  "destination station code"
// @Param        Date 			body	string  true  "date of booking"
// @Success      200  {array} 	models.Train
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /train/searchRoute [post]
func SearchRoute() gin.HandlerFunc {
	return func(g *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()

		var search SearchQuery

		if err := g.BindJSON(&search); err != nil {
			g.JSON(http.StatusOK, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("fail in gin json bind")
			return
		}

		log.Debug("gin json binding done.")

		cursor, err := trainCollection.Find(
			ctx,
			bson.D{{"Stations", bson.D{{"$all", bson.A{search.Source, search.Destination}}}}},
		)

		if err != nil {
			g.JSON(http.StatusOK, gin.H{"error": "no routes"})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("no routes present")
			return
		}

		var trainDetails []models.Train

		if err := cursor.All(ctx, &trainDetails); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).
				Error("error in response cursor of mongo db")
			return
		}

		//-------------check for date availability
		layout := "01/02/06"
		parseDate, err := time.Parse(layout, search.Date)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("fail to parse search date")
			return
		}

		if time.Now().After(parseDate) {
			g.JSON(http.StatusBadRequest, gin.H{"error": "invalid date"})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("invalid date")
			return
		}

		weekday := parseDate.Weekday().String()
		responseTrainDetails := helper.FilterDetailsOnWeekdayAvailability(weekday, trainDetails)
		log.Debug("filtering train details on the basis of week day availability is done")
		responseTrainDetails = helper.CalculatePriceMany(
			search.Source,
			search.Destination,
			responseTrainDetails,
		)

		g.JSON(http.StatusOK, gin.H{"trainDetails": responseTrainDetails})
		log.Info("succesfully sending train details")
	}
}

// ShowAccount godoc
// @Summary      Get details of sepecific train with its number
// @Description  Get details of sepecific train with its number
// @Tags         Train
// @Accept       json
// @Produce      json
// @Param        TrainNumber  body 	string  true  "unique train number"
// @Success      200  {object} 	models.Train
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /train/trainDetails [post]
func TrainDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
		defer cancel()

		var body struct {
			TrainNumber string
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("fail in gin json bind")
			return
		}
		log.Debug("gin json binding done.")

		var TrainDetails models.Train
		err := trainCollection.FindOne(ctx, bson.M{"TrainNumber": body.TrainNumber}).
			Decode(&TrainDetails)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).
				Error("fail to fetch records from DB")
			return
		}

		log.WithFields(logrus.Fields{"TrainNumber": body.TrainNumber}).
			Debug("found train record of given train number.")
		c.JSON(http.StatusOK, gin.H{"trainDetails": TrainDetails})
	}
}
