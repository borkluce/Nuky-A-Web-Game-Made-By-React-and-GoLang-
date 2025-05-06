import { create } from "zustand"

// Custom Axios instance with preset base URL
import { CAxios } from "../../core/configs/cAxios"

// Types
import {
    CooldownLeftInSecondsRequest,
    CooldownLeftInSecondsResponse,
    LoginRequest,
    LoginResponse,
    RegisterRequest,
    RegisterResponse,
} from "../types/user.dtos"
import { User } from "../types/user"

interface useUserState {
    // Fields
    user: User | null
    coolDate: number // Datetime cooldown will be finished

    // APIs
    login: (LoginRequest: LoginRequest) => Promise<LoginResponse>
    register: (registerRequest: RegisterRequest) => Promise<RegisterResponse>
    cooldDownLefInSeconds: (
        cooldownLeftInSecondsRequest: CooldownLeftInSecondsRequest
    ) => Promise<CooldownLeftInSecondsResponse>

    // Behaviours
    isAllowedToMove: () => boolean
}

export const useUser = create<useUserState>((set, get) => ({
    user: null,
}))
