package model

type Province struct {
	ID               int    `json:"id"`
	ProvinceName     string `json:"province_name"`
	ProvinceColorHex string `json:"province_color_hex"`
	AttackCount      int    `json:"attack_count"`
	SupportCount     int    `json:"support_count"`
}
