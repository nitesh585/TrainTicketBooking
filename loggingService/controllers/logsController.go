package controllers

import (
	"context"
	"fmt"
	"loggingService/database"

	"go.mongodb.org/mongo-driver/mongo"
)

func SaveInDB(ctx context.Context, msg interface{}, topic string) {
	var logsCollection *mongo.Collection = database.OpenCollection(database.MongoClient, topic)
	insertNumber, err := logsCollection.InsertOne(ctx, msg)

	if err != nil {
		return
	}
	fmt.Printf("Log saved! insertNumber(%d)  ", insertNumber)
}
