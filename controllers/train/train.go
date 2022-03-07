package controller

import (
	"context"
	"net/http"
	"rail/database"
	models "rail/models/train"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SearchQuery struct {
	Source      string
	Destination string
	Date        string
}

var userCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "train")

func SearchRoute() gin.HandlerFunc {
	return func(g *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()

		var search SearchQuery

		if err := g.BindJSON(&search); err != nil {
			g.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}

		var trainDetails models.Train

		err := userCollection.FindOne(ctx, bson.D{{"tags", bson.D{{"$all", bson.A{"red", "blank"}}}}}).Decode(&trainDetails)
		if err != nil {
			g.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}

	}
}
