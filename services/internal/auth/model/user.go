package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `json:"ID" bson:"_id,omitempty"`

	Username     string    `json:"username" bson:"username"`
	Email        string    `json:"email" bson:"email"`
	LastMoveDate time.Time `json:"last_move_date" bson:"lastMoveDate"`

	Password string `bson:"password"`
}
