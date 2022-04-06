package controllers

import (
	"emailService/logger"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var log logrus.Logger = *logger.GetLogger()
var FROM string
var PASSWORD string
var HOST string
var SMTP_DEFAULT_PORT string

func init() {
	// initialize all the constants
	err := godotenv.Load(".env")
	if err != nil {
		log.WithFields(logrus.Fields{"err": err.Error()}).Error("Failed to load .env file")
		return
	}

	FROM = os.Getenv("FROM")
	PASSWORD = os.Getenv("PASSWORD")
	HOST = os.Getenv("HOST_URL")
	SMTP_DEFAULT_PORT = os.Getenv("SMTP_DEFAULT_PORT")
}

func SendEmail(receiverEmail, userName string) {
	toList := []string{receiverEmail}

	msg := "Welcome " + userName + "! to rail services."
	body := []byte(msg)
	auth := smtp.PlainAuth("", FROM, PASSWORD, HOST)

	err := smtp.SendMail(HOST+":"+SMTP_DEFAULT_PORT, auth, FROM, toList, body)
	if err != nil {
		log.WithFields(logrus.Fields{"err": err.Error()}).Error("Not able to send email")
	}

	log.WithFields(logrus.Fields{"userName": userName, "email": receiverEmail}).Info("Mail sent!")
}
