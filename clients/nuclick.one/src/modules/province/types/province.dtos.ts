import { Province } from "./province"
// --------------------------------------------------------------------
export type GetAllProvincesRequest = {}
export type GetAllProvnceseResponse = {
    provinceList: Province[]
}
// --------------------------------------------------------------------
export type AttackProvinceRequest = {
    provinceID: string
}
export type AttackProvinceResponse = {
    isSuccess: boolean
}
// --------------------------------------------------------------------
export type SupportProvinceRequest = {
    provinceID: string
}
export type SupportProvinceResponse = {
    isSuccess: boolean
}
