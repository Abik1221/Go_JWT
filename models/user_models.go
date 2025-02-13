package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id" omitEmpty:"true"`
	First_Name    string             `json:"first_name" validate:"required"`
	Last_Name     string             `json:"last_name" validate:"required"`
	Email         string             `json:"email" validate:"requires"`
	Phone         string             `json:"phone" validate:"required"`
	Token         string             `json:"token"`
	Refresh_Token string             `json:"refresh_token"`
	User_Type     string             `json:"user_type"`
	Created_At    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"Updated_At"`
	User_Id       string             `json:"user_id"`
}
