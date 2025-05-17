import { create } from "zustand"

// Custom Axios instance with preset base URL
import { CAxios } from "../../core/configs/cAxios"

// Types
import { Province } from "../types/province"

// Dtos
import {
    GetAllProvincesRequest,
    GetAllProvincesResponse,
    GetTopProvincesResponse,
    AttackProvinceRequest,
    AttackProvinceResponse,
    SupportProvinceRequest,
    SupportProvinceResponse
} from "../types/province.dtos"

interface useProvinceState {
    provinceList: Province[] | null
    topProvinces: Province[] | null
    getAllProvinces: (
        request: GetAllProvincesRequest
    ) => Promise<GetAllProvincesResponse>
    getTopProvinces: () => Promise<GetTopProvincesResponse>
    attackProvince: (
        request: AttackProvinceRequest
    ) => Promise<AttackProvinceResponse>
    supportProvince: (
        request: SupportProvinceRequest
    ) => Promise<SupportProvinceResponse>
}

export const useProvince = create<useProvinceState>((set) => ({
    provinceList: null,
    topProvinces: null,
    
    getAllProvinces: async (request: GetAllProvincesRequest) => {
        const response = await CAxios.get<GetAllProvincesResponse>('/provinces', { params: request });
        // Handle both possible response structures
        let provinces: Province[] = [];
        if (Array.isArray(response.data)) {
            // If the response is directly an array
            provinces = response.data;
        } else if (response.data.province_list && Array.isArray(response.data.province_list)) {
            // If the response has a provinces property that is an array
            provinces = response.data.province_list;
        }
        set({ provinceList: provinces });
        return response.data;
    },
    
    getTopProvinces: async () => {
        const response = await CAxios.get<GetTopProvincesResponse>('/provinces/top');
        // Handle both possible response structures
        let provinces: Province[] = [];
        if (Array.isArray(response.data)) {
            // If the response is directly an array
            provinces = response.data;
        } else if (response.data.provinces && Array.isArray(response.data.provinces)) {
            // If the response has a provinces property that is an array
            provinces = response.data.provinces;
        }
        set({ topProvinces: provinces });
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
