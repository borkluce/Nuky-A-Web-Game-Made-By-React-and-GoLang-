import { create } from "zustand"
import { persist } from "zustand/middleware"

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
import { User, createUserFromResponse } from "../types/user"

const DEFAULT_COOLDOWN_SECONDS = 3600 // 1 hour cooldown

interface useUserInterface {
    // Fields
    user: User | null
    coolDate: number // Datetime cooldown will be finished

    // APIs
    login: (loginRequest: LoginRequest) => Promise<boolean>
    register: (registerRequest: RegisterRequest) => Promise<boolean>
    cooldownLeftInSeconds: (
        cooldownLeftInSecondsRequest: CooldownLeftInSecondsRequest
    ) => Promise<CooldownLeftInSecondsResponse>
    logout: () => void

    // Move related
    resetCooldownAfterMove: (seconds?: number) => void

    // Behaviours
    isAllowedToMove: () => boolean
    updateCooldown: (seconds: number) => void
    getRemainingCooldownSeconds: () => number
    isAuthenticated: () => boolean
}

export const useUser = create<useUserInterface>()(
    persist(
        (set, get) => ({
            user: null,
            coolDate: 0,
            // --------------------------------------------------------------------
            login: async (loginRequest: LoginRequest) => {
                try {
                    const response = await CAxios.post<LoginResponse>(
                        "/auth/login",
                        loginRequest
                    )

                    if (response.data.token) {
                        let user: User
                        
                        if (response.data.user) {
                            // Use user data from response
                            user = createUserFromResponse(response.data.token, response.data.user)
                        } else {
                            // Fallback to basic user creation
                            user = createUserFromResponse(response.data.token, {
                                username: "",
                                email: loginRequest.email || ""
                            })
                        }

                        set({ user })
                        return true
                    }
                    return false
                } catch (error) {
                    console.error("Login failed:", error)
                    throw error
                }
            },
            // --------------------------------------------------------------------
            register: async (registerRequest: RegisterRequest) => {
                try {
                    const response = await CAxios.post<RegisterResponse>(
                        "/auth/register",
                        registerRequest
                    )

                    if (response.data.token) {
                        let user: User
                        
                        if (response.data.user) {
                            // Use user data from response
                            user = createUserFromResponse(response.data.token, response.data.user)
                        } else {
                            // Fallback to basic user creation
                            user = createUserFromResponse(response.data.token, {
                                username: registerRequest.username,
                                email: registerRequest.email
                            })
                        }

                        set({ user })
                        return true
                    }
                    return false
                } catch (error) {
                    console.error("Registration failed:", error)
                    throw error
                }
            },
            // --------------------------------------------------------------------
            cooldownLeftInSeconds: async (
                cooldownRequest: CooldownLeftInSecondsRequest
            ) => {
                try {
                    const response =
                        await CAxios.post<CooldownLeftInSecondsResponse>(
                            "/user/cooldown",
                            cooldownRequest
                        )

                    const cooldownSeconds =
                        response.data.cooldown_left_in_seconds
                    get().updateCooldown(cooldownSeconds)

                    return response.data
                } catch (error) {
                    console.error("Failed to fetch cooldown:", error)
                    throw error
                }
            },
            // --------------------------------------------------------------------
            logout: () => {
                set({ user: null, coolDate: 0 })
            },
            // --------------------------------------------------------------------
            isAllowedToMove: () => {
                const { coolDate } = get()
                return Date.now() >= coolDate
            },
            // --------------------------------------------------------------------
            updateCooldown: (seconds: number) => {
                const newCoolDate = Date.now() + seconds * 1000
                set({ coolDate: newCoolDate })

                if (get().user) {
                    set({
                        user: {
                            ...get().user!,
                            last_move_date: new Date(),
                        },
                    })
                }
            },
            // --------------------------------------------------------------------
            getRemainingCooldownSeconds: () => {
                const { coolDate } = get()
                const remaining = Math.max(0, (coolDate - Date.now()) / 1000)
                return Math.ceil(remaining)
            },
            // --------------------------------------------------------------------
            isAuthenticated: () => {
                const { user } = get()
                return !!(user?.token && user?.isAuthenticated)
            },
            // --------------------------------------------------------------------
            resetCooldownAfterMove: async (seconds?: number) => {
                // If seconds are provided, use them
                // Otherwise, use default or request from server
                if (seconds !== undefined) {
                    get().updateCooldown(seconds)
                    return
                }

                try {
                    // Option 1: Use a fixed cooldown
                    // This is temporary. I will fix after backend is finished
                    get().updateCooldown(DEFAULT_COOLDOWN_SECONDS)

                    // Option 2: Request cooldown from server after move
                    // const token = get().user?.token
                    // if (token) {
                    //    await get().cooldownLeftInSeconds({ token })
                    // }

                    // Update last_move_date
                    if (get().user) {
                        set({
                            user: {
                                ...get().user!,
                                last_move_date: new Date(),
                            },
                        })
                    }
                } catch (error) {
                    console.error("Failed to reset cooldown:", error)
                    // Fallback to default cooldown if server request fails
                    get().updateCooldown(DEFAULT_COOLDOWN_SECONDS)
                }
            },
        }),
        {
            name: "user",
            partialize: (state) => ({
                user: state.user,
                coolDate: state.coolDate,
            }),
        }
    )
)