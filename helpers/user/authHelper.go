package helper

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	uuid      uuid.UUID
	jwt.StandardClaims
}

var SECRET_KEY = os.Getenv("SECRET_KEY")

func VerifyToken() {

}

func CreateToken(email, first_name, last_name string, uuid uuid.UUID) (string, error) {
	claims := SignedDetails{
		Email:     email,
		FirstName: first_name,
		LastName:  last_name,
		uuid:      uuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return "", err
	}

	return token, err
}
