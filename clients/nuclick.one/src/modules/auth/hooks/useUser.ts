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
import { User } from "../types/user"

const DEFAULT_COOLDOWN_SECONDS = 3600 // 1 hour cooldown

interface useUserInterface {
    // Fields
    user: User | null
    coolDate: number // Datetime cooldown will be finished

    // APIs
    login: (loginRequest: LoginRequest) => Promise<LoginResponse>
    register: (registerRequest: RegisterRequest) => Promise<RegisterResponse>
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
}

export const useUser = create<useUserInterface>()(
    persist(
        (set, get) => ({
            user: null,
            coolDate: 0,

            login: async (loginRequest: LoginRequest) => {
                try {
                    const response = await CAxios.post<LoginResponse>(
                        "/auth/login",
                        loginRequest
                    )
                    const { user } = response.data

                    set({
                        user: {
                            username: user.username,
                            email: user.email,
                            lastMoveDate: new Date(user.lastMoveDate),
                        },
                    })

                    return response.data
                } catch (error) {
                    console.error("Login failed:", error)
                    throw error
                }
            },

            register: async (registerRequest: RegisterRequest) => {
                try {
                    const response = await CAxios.post<RegisterResponse>(
                        "/auth/register",
                        registerRequest
                    )
                    const { user } = response.data

                    set({
                        user: {
                            username: user.username,
                            email: user.email,
                            lastMoveDate: new Date(user.lastMoveDate),
                        },
                    })

                    return response.data
                } catch (error) {
                    console.error("Registration failed:", error)
                    throw error
                }
            },

            cooldownLeftInSeconds: async (
                cooldownRequest: CooldownLeftInSecondsRequest
            ) => {
                try {
                    const response =
                        await CAxios.post<CooldownLeftInSecondsResponse>(
                            "/user/cooldown",
                            cooldownRequest
                        )

                    const cooldownSeconds = response.data.cooldownLeftInSeconds
                    get().updateCooldown(cooldownSeconds)

                    return response.data
                } catch (error) {
                    console.error("Failed to fetch cooldown:", error)
                    throw error
                }
            },

            logout: () => {
                set({ user: null, coolDate: 0 })
            },

            isAllowedToMove: () => {
                const { coolDate } = get()
                return Date.now() >= coolDate
            },

            updateCooldown: (seconds: number) => {
                const newCoolDate = Date.now() + seconds * 1000
                set({ coolDate: newCoolDate })

                if (get().user) {
                    set({
                        user: {
                            ...get().user!,
                            lastMoveDate: new Date(),
                        },
                    })
                }
            },

            getRemainingCooldownSeconds: () => {
                const { coolDate } = get()
                const remaining = Math.max(0, (coolDate - Date.now()) / 1000)
                return Math.ceil(remaining)
            },

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

                    // Update lastMoveDate
                    if (get().user) {
                        set({
                            user: {
                                ...get().user!,
                                lastMoveDate: new Date(),
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
