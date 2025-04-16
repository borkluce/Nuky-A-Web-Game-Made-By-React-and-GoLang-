package model

import "time"

type User struct {
	ID           int
	Email        string
	LastMoveDate time.Time
}
