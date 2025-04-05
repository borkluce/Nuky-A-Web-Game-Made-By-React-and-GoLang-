package model

import "gorm.io/gorm"

type State struct {
	gorm.Model
	ID            int    `gorm:"primaryKey;autoIncrement"`
	StateName     string `gorm:"size:255;not null"`
	StateColorHex string `gorm:"size:7;not null"` // Hex renk kodu için max 7 karakter (#FFFFFF)
	AttackCount   int    `gorm:"default:0"`
	SupportCount  int    `gorm:"default:0"`
}
