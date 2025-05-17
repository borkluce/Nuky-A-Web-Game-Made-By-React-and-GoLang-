import { Province } from "./province"

// --------------------------------------------------------------------

export type GetAllProvincesRequest = {}

export type GetAllProvincesResponse = {
    province_list: Province[]
}
// --------------------------------------------------------------------

export type GetTopProvincesRequest = {}

export type GetTopProvincesResponse = {
    provinces?: Province[]
}

// --------------------------------------------------------------------

export type AttackProvinceRequest = {
    province_ID: string
}

export type AttackProvinceResponse = {
    is_success: boolean
}
// --------------------------------------------------------------------

export type SupportProvinceRequest = {
    province_ID: string
}

export type SupportProvinceResponse = {
    is_success: boolean
}
