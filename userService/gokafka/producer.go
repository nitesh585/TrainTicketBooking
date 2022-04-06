package gokafka

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"
	"userService/logger"

	kafka "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var brokers = []string{
	"localhost:9091",
}
var log logrus.Logger = *logger.GetLogger()

func getRandomKey() []byte {
	var src = rand.NewSource(time.Now().UnixNano())
	return []byte(string(src.Int63()))
}

func getProducer(ctx context.Context, topic string, brokers []string) *kafka.Writer {
	// intialize the writer with the broker addresses, and the topic
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})
}

func WriteMsgToKafka(topic string, msg interface{}) (bool, error) {
	ctx := context.Background()
	jsonString, err := json.Marshal(msg)
	msgString := string(jsonString)

	if err != nil {
		return false, err
	}

	writer := getProducer(ctx, topic, brokers)
	defer writer.Close()

	err = writer.WriteMessages(ctx, kafka.Message{
		Key:   getRandomKey(),
		Value: []byte(msgString),
	})

	if err != nil {
		log.WithFields(logrus.Fields{"error": err.Error(), "msg": msg}).
			Error("could not write message ")
		return false, err
	}

	return true, nil
}
