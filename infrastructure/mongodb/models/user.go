package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username,omitempty"`
	Password  string             `json:"password" bson:"password,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at,omitempty"`
}
