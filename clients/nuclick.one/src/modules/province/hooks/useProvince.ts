import { create } from "zustand"

// Custom Axios instance with preset base URL
import { CAxios } from "../../core/configs/cAxios"

// Types
import { Province } from "../types/province"

interface useProvinceState {
    province_list: Province[] | null
    getAllProvinces: () => Promise<>
}
