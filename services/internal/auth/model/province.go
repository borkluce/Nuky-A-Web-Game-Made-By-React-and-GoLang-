package model

import "gorm.io/gorm"

type Province struct {
	gorm.Model
	ID               string `gorm:"primaryKey"`
	ProvinceName     string `gorm:"size:255;not null"`
	ProvinceColorHex string `gorm:"size:7;not null"` // Hex renk kodu için max 7 karakter (#FFFFFF)
	AttackCount      int    `gorm:"default:0"`
	SupportCount     int    `gorm:"default:0"`
}
