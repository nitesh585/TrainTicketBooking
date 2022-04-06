package helper

import (
	"crypto/rand"
	"encoding/hex"
	"paymentService/logger"

	"github.com/sirupsen/logrus"
)

var log logrus.Logger = *logger.GetLogger()

func RandRefrenceID(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		log.WithFields(logrus.Fields{"err": err.Error()}).Error("fail to random read bytes")

		return "", err
	}

	log.Debug("rand refernce id is created")
	return hex.EncodeToString(bytes), nil
}
