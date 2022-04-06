package helper

import (
	"errors"
	"os"
	"time"
	"userService/logger"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

var log logrus.Logger = *logger.GetLogger()

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	User_id   string
	jwt.StandardClaims
}

var SECRET_KEY = os.Getenv("SECRET_KEY")

func VerifyToken(clientToken string) (claims *SignedDetails, responseErr error) {
	token, err := jwt.ParseWithClaims(
		clientToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		responseErr = err
		log.WithFields(logrus.Fields{"error": err.Error()}).Error("jwt parsing with claims failed.")
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		responseErr = errors.New("invalid token")
		log.WithFields(logrus.Fields{"error": err.Error()}).Error("invalid token")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		responseErr = errors.New("token is expired")
		log.WithFields(logrus.Fields{"error": err.Error()}).Error("token expired")
		return
	}

	log.Debug("token verified")
	return claims, responseErr
}

func CreateToken(email, first_name, last_name, user_id string) (string, error) {
	claims := SignedDetails{
		Email:     email,
		FirstName: first_name,
		LastName:  last_name,
		User_id:   user_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.WithFields(logrus.Fields{"error": err.Error()}).Error("jwt new claims failed")
		return "", err
	}

	log.Debug("successfully created token")
	return token, err
}
