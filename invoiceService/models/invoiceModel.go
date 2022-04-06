package models

import (
	"invoiceService/logger"
	"time"

	"github.com/sirupsen/logrus"
)

var log logrus.Logger = *logger.GetLogger()

type Invoice struct {
	ID            string
	UserID        string
	BookingID     string
	TransactionID string
	CreatedAt     time.Time
}

type invoiceMethod interface {
	SendEmailInvoice()
	SendMobileInvoice()
}

func (invoice *Invoice) SendEmailInvoice() {
	log.Info("Email invoice sent!")
}

func (invoice *Invoice) SendSMSInvoice() {
	log.Info("Mobile invoice sent!")
}
