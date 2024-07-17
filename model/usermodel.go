package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Email    string             `json:"email" validate:"email,required"`
	Password string             `json:"password" validate:"required"`
	Username string             `json:"username"`
	City     string             `json:"city"`
	State    string             `json:"state"`
	Country  string             `json:"country"`
	Age      int                `json:"age"`
	Token    string             `json:"token"`
	PinCode  int                `json:"pinCode"`
}
