package gokafka

import (
	"context"
	"encoding/json"
	"invoiceService/logger"
	"invoiceService/models"

	kafka "github.com/segmentio/kafka-go"
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

func InvoiceConsumer() {
	for {
		ctx := context.Background()

		// the `ReadMessage` method blocks until we receive the next event
		reader := getKafkaReader(ctx, "invoice", "invoice-id", []string{"localhost:9091"})
		msg, err := reader.ReadMessage(ctx)
		var invoice models.Invoice

		json.Unmarshal(msg.Value, &invoice)

		if err != nil {
			panic("could not read message " + err.Error())
		}

		log.WithFields(logrus.Fields{
			"inv": invoice,
		}).Info("invoice fetched")

		go invoice.SendEmailInvoice()
		go invoice.SendSMSInvoice()
		log.Info("invoice sent both via email and sms.")
	}
}
