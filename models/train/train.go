package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Train struct {
	Id              primitive.ObjectID `json:"_id"`
	TrainName       string             `json:"trainname"`
	TrainNumber     string             `json:"trainNumber"`
	FromStationCode string             `json:"source"`
	ToStationCode   string             `json:"destination"`
	ArrivalTime     string             `json:"avlTime"`
	DepartureTime   string             `json:"depTime"`
	Distance        string             `json:"distance"`
	Duration        string             `json:"duration"`
	RunningMon      string             `json:"runningMon"`
	RunningTue      string             `json:"runningTue"`
	RunningWed      string             `json:"runningWed"`
	RunningThr      string             `json:"runningThr"`
	RunningFri      string             `json:"runningFri"`
	RunningSat      string             `json:"runningSat"`
	RunningSun      string             `json:"runningSun"`
	AvlClasses      []string           `json:"avlClasses"`
	Stations        []string           `json:"stations"`
	Price           []int
}
