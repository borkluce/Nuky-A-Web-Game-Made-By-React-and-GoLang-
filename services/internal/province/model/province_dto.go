package model

// --------------------------------------------------------------------

type GetAllProvinceResponse struct {
	ProvinceList []Province `json:"province_list"`
}

// --------------------------------------------------------------------

type AttackProvinceRequest struct {
	ProvinceID string `json:"province_id"`
}

type AttackProvinceResponse struct {
	IsSuccess bool `json:"is_success"`
}

// --------------------------------------------------------------------

type SupportProvinceRequest struct {
	ProvinceID string `json:"province_id"`
}

type SupportProvinceResponse struct {
	IsSuccess bool `json:"is_success"`
}
