package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	provmodel "services/internal/province/model"
)

type Game struct {
	ID             primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	LoserProvinces []provmodel.Province `json:"loserProvinces" bson:"loserProvinces"`
	UpdatedDate    time.Time            `json:"updatedDate" bson:"updatedDate"`
	DeletedDate    *time.Time           `json:"deletedDate,omitempty" bson:"deletedDate,omitempty"`
}
