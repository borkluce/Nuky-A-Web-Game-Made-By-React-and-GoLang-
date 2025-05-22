package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Province struct {
	ID               primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	ProvinceName     string             `json:"province_name" bson:"provinceName"`
	ProvinceColorHex string             `json:"province_color_hex" bson:"provinceColorHex"`
	AttackCount      int                `json:"attack_count" bson:"attackCount"`
	SupportCount     int                `json:"support_count" bson:"supportCount"`
	DestroymentRound int                `json:"destroyment_round" bson:"destroymentRound"`
}

func (p Province) MongoIDToStringID(mongoID primitive.ObjectID) (string, error) {
	stringedID := mongoID.Hex()
	return stringedID, nil
}
