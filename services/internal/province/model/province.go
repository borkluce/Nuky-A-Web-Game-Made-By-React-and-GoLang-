package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Province struct {
	ID               primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	ProvinceName     string             `json:"provinceName"`
	ProvinceColorHex string             `json:"provinceColorHex"`
	AttackCount      int                `json:"attackCount"`
	SupportCount     int                `json:"supportCount"`
}

func (p Province) MongoIDToStringID(mongoID primitive.ObjectID) (string, error) {
	stringedID := mongoID.Hex()
	return stringedID, nil
}
