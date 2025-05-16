import { Province } from "./province"

// --------------------------------------------------------------------

export type GetAllProvincesRequest = {}

export type GetAllProvnceseResponse = {
    province_list: Province[]
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
