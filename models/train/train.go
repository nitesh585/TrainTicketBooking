package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TrainDetails struct {
	Id          primitive.ObjectID `bson:"_id"`
	TrainNumber int16              `json:"trainNumber"`
	ArrivalTime time.Time          `json:"arrivalTime"`
	Departure   time.Time          `json:"departureTime"`
}

type Station struct {
	Id           primitive.ObjectID `json:"_id"`
	Name         string             `json:"name"`
	TrainDetails TrainDetails       `json:"trainDetails"`
}

type Train struct {
	Id          primitive.ObjectID `json:"_id"`
	Name        string             `json:"name"`
	TrainNumber string             `json:"trainNumber"`
	Source      string             `json:"source"`
	Destination string             `json:"destination"`
	Stations    []Station          `json:"stations"`
}
