import { create } from "zustand"

// Custom Axios instance with preset base URL
import { CAxios } from "../../core/configs/cAxios"

// Types
import { Province } from "../types/province"

// Dtos
import {
    GetAllProvincesRequest,
    GetAllProvnceseResponse,
    AttackProvinceRequest,
    AttackProvinceResponse,
    SupportProvinceRequest,
    SupportProvinceResponse
} from "../types/province.dtos"

interface useProvinceState {
    provinceList: Province[] | null
    getAllProvinces: (
        request: GetAllProvincesRequest
    ) => Promise<GetAllProvnceseResponse>
    attackProvince: (
        request: AttackProvinceRequest
    ) => Promise<AttackProvinceResponse>
    supportProvince: (
        request: SupportProvinceRequest
    ) => Promise<SupportProvinceResponse>
}

export const useProvince = create<useProvinceState>(() => ({
    provinceList: null,
    
    getAllProvinces: async (request: GetAllProvincesRequest) => {
        const response = await CAxios.get<GetAllProvnceseResponse>('/provinces', { params: request });
        return response.data;
    },
    
    attackProvince: async (request: AttackProvinceRequest) => {
        const response = await CAxios.post<AttackProvinceResponse>('/provinces/attack', request);
        return response.data;
    },
    
    supportProvince: async (request: SupportProvinceRequest) => {
        const response = await CAxios.post<SupportProvinceResponse>('/provinces/support', request);
        return response.data;
    }
}))