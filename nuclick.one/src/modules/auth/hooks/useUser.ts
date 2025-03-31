import { create } from "zustand"
import axios from "axios"

// Types
import { LoginRequest, LoginResponse } from "../types/user.dtos"
import { User } from "../types/user"

interface useUserState {
    user: User | null
    login: (loginRequest: LoginRequest) => Promise<LoginResponse>
    // register: (registerRequest: RegisterRequest) => Promise<RegisterResponse>
}

export const useUser = create<useUserState>((set, get) => ({
    user: null,
    login: async (loginRequest) => {
        const response = await axios.post<LoginResponse>(
            "/auth/login",
            loginRequest
        )

        return response.data
    },
}))
