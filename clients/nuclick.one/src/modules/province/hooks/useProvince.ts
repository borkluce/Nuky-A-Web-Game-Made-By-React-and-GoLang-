import { create } from "zustand"

// Custom Axios instance
import { CAxios } from "../../core/configs/cAxios"

// Types
import { generateRandomProvinces, Province } from "../types/province"

// DTOs
import {
    GetAllProvinceResponse,
    GetTopProvincesResponse,
    AttackProvinceRequest,
    AttackProvinceResponse,
    SupportProvinceRequest,
    SupportProvinceResponse,
} from "../types/province.dtos"

interface useProvinceState {
    // States
    provinceList: Province[]
    topProvinces: Province[]

    // Status
    isLoading: boolean
    error: string | null

    // APIs
    getAllProvinces: () => Promise<void>
    getTopProvinces: () => Promise<void>
    attackProvince: (
        request: AttackProvinceRequest
    ) => Promise<AttackProvinceResponse>
    supportProvince: (
        request: SupportProvinceRequest
    ) => Promise<SupportProvinceResponse>
}

export const useProvince = create<useProvinceState>((set) => ({
    provinceList: [],
    topProvinces: [],
    isLoading: false,
    error: null,
    // --------------------------------------------------------------------
    getAllProvinces: async () => {
        set({ isLoading: true, error: null })
        try {
            const response = await CAxios.get<GetAllProvinceResponse>(
                //    "/province?type=all"
                "/province"
            )

            console.log(response.data)

            set({
                provinceList: response.data.province_list,
                isLoading: false,
            })
        } catch (error) {
            console.error("Failed to fetch provinces", error)
            set({ error: "Failed to load provinces", isLoading: false })
        }
    },
    // -------------------------------------------------------------------
    getTopProvinces: async () => {
        set({ isLoading: true, error: null })
        try {
            const response = await CAxios.get<GetTopProvincesResponse>(
                "/province/top"
            )
            console.log("top 5 states")
            console.log(response.data.provinces)

            set({
                topProvinces: response.data.provinces,
                isLoading: false,
            })
        } catch (error) {
            console.error("Failed to fetch provinces", error)
            set({ error: "Failed to load provinces", isLoading: false })
        }
    },
    // --------------------------------------------------------------------
    attackProvince: async (request: AttackProvinceRequest) => {
        const response = await CAxios.post<AttackProvinceResponse>(
            "/province/attack",
            request
        )
        return response.data
    },
    // --------------------------------------------------------------------
    supportProvince: async (request: SupportProvinceRequest) => {
        const response = await CAxios.post<SupportProvinceResponse>(
            "/province/support",
            request
        )
        return response.data
    },
}))
