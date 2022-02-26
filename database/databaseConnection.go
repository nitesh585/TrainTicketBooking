package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBintance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
	MongoDb_URL := os.Getenv("MONGO_DB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb_URL))

	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Mongo DB")

	return client
}

var MongoClient *mongo.Client = DBintance()

func OpenCollection(client *mongo.Client, colletionName string) *mongo.Collection {
	// TODO: write DatabaseName
	var collection *mongo.Collection = client.Database("trainTicket").Collection(colletionName)
	return collection
}
