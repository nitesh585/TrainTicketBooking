package database

import (
	"context"
	"os"
	"time"
	"trainService/logger"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var log logrus.Logger = *logger.GetLogger()

func DBintance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.WithFields(logrus.Fields{"error": err.Error()}).Error("failed to load .env file")
	}
	MongoDb_URL := os.Getenv("MONGO_DB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb_URL))

	if err != nil {
		log.WithFields(logrus.Fields{"error": err.Error()}).
			Error("new mongo client creation failed")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.WithFields(logrus.Fields{"error": err.Error()}).Error("client connection failed")
	}

	log.Info("Connected to Mongo DB")
	log.Trace("Connected to Mongo DB")
	return client
}

var MongoClient *mongo.Client = DBintance()

func OpenCollection(client *mongo.Client, colletionName string) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.WithFields(logrus.Fields{"err": err.Error()}).Error("Failed to load .env file")
	}

	DATABASE := os.Getenv("DATABASE")
	var collection *mongo.Collection = client.Database(DATABASE).Collection(colletionName)
	return collection
}
