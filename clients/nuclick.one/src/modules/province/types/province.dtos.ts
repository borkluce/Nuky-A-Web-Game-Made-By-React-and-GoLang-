import { Province } from "./province"

// --------------------------------------------------------------------

export type GetAllProvinceResponse = {
    province_list: Province[]
}
// --------------------------------------------------------------------

export type GetTopProvincesRequest = {}

export type GetTopProvincesResponse = {
    provinces?: Province[]
}

// --------------------------------------------------------------------

export type AttackProvinceRequest = {
    province_id: string
}

export type SupportProvinceRequest = {
    province_id: string
}

// --------------------------------------------------------------------

export type AttackProvinceResponse = {
    is_success: boolean
}

export type SupportProvinceResponse = {
    is_success: boolean
}
