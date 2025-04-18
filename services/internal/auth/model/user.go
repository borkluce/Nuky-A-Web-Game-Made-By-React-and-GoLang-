package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `bson:"password"`
	LastMoveDate time.Time          `json:"lastMoveDate" bson:"lastMoveDate"`
}
