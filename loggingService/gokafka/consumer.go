package gokafka

import (
	"context"
	"encoding/json"
	"loggingService/controllers"
	"loggingService/logger"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var log logrus.Logger = *logger.GetLogger()

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

func LogConsumer(ctx context.Context, topic string) {
	reader := getKafkaReader(ctx, topic, "logs-id", []string{"localhost:9092"})
	msg, err := reader.ReadMessage(ctx)

	var logMsg interface{}

	json.Unmarshal(msg.Value, &logMsg)
	if err != nil {
		panic("could not read message " + err.Error())
	}

	log.WithFields(logrus.Fields{
		"logMsg": logMsg,
	}).Info("Log message fetched")

	go controllers.SaveInDB(ctx, msg, topic)
}
