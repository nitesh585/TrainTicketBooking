package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	//TODO: add it to model and use "bson inline" during unmarshal
	Street  string
	City    string
	Country string
}

type User struct {
	Id         primitive.ObjectID `bson:"_id"`
	FirstName  string             `json:"first_name" validate:"required, min=2, max=100"`
	LastName   string             `json:"last_name" validate:"required, min=2, max=100"`
	Password   string             `json:"password" validate:"required, min=6"`
	Email      string             `json:"email" validate:"email, required"`
	Phone      string             `json:"phone" validate:"required"`
	Token      string             `json:"token"`
	Created_at time.Time          `json:"created_at"`
	User_id    string             `json:"user_id" `
}
