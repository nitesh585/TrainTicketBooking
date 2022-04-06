package gokafka

import (
	"context"
	"emailService/controllers"
	"emailService/logger"
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var log logrus.Logger = *logger.GetLogger()

var EMAIL string
var GROUP_ID string
var KAFKA_BROKER string

func init() {
	// initialize all the constants
	err := godotenv.Load(".env")
	if err != nil {
		log.WithFields(logrus.Fields{"err": err.Error()}).Error("Failed to load .env file")
		return
	}

	KAFKA_BROKER = os.Getenv("FROM")
	EMAIL = os.Getenv("FROM")
	GROUP_ID = os.Getenv("FROM")

}

func getKafkaReader(ctx context.Context, topic, groupID string, brokers []string) *kafka.Reader {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})
	return r
}

func EmailConsumer() {
	for {
		ctx := context.Background()

		// the `ReadMessage` method blocks until we receive the next event
		reader := getKafkaReader(ctx, EMAIL, GROUP_ID, []string{KAFKA_BROKER})
		msg, err := reader.ReadMessage(ctx)
		var user struct {
			Email string
			Name  string
		}

		json.Unmarshal(msg.Value, &user)

		if err != nil {
			log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Info("not able to read message")
		} else {
			log.WithFields(logrus.Fields{
				"user": user,
			}).Info("User details fetched")

			go controllers.SendEmail(user.Email, user.Name)
		}
	}
}
