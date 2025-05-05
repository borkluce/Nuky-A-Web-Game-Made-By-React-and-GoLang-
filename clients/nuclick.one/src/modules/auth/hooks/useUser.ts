import { create } from "zustand"

// Custom Axios instance with preset base URL
import { CAxios } from "../../core/configs/cAxios"

// Types
import {
    LoginRequest,
    LoginResponse,
    RegisterRequest,
    RegisterResponse,
} from "../types/user.dtos"
import { User } from "../types/user"

interface useUserState {
    user: User | null
    // login: (loginRequest: LoginRequest) => Promise<void>
    register: (registerRequest: RegisterRequest) => Promise<void>
    isAllowedToMove: () => boolean
}

export const useUser = create<useUserState>((set, get) => ({
    user: null,
    // login: async (loginRequest) => {
    //     const response = await CAxios.post<LoginResponse>(
    //         "/auth/login",
    //         loginRequest
    //     )
    //     set({ user: get({ user }) })
    // },
    register: async (registerRequest) => {
        const response = await CAxios.post<RegisterRequest>(
            "/auth/register",
            registerRequest
        )
    },
    isAllowedToMove: () => {
        return true
    },
}))
