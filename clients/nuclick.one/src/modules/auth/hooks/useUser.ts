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
import { User, createEmptyUser } from "../types/user"

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
    
    // Behaviours
    isAllowedToMove: () => boolean
    updateCooldown: (seconds: number) => void
}

export const useUser = create<useUserInterface>()(
    persist(
        (set, get) => ({
            user: null,
            coolDate: 0,
            
            login: async (loginRequest: LoginRequest) => {
                try {
                    const response = await CAxios.post<LoginResponse>('/auth/login', loginRequest)
                    const { token, user } = response.data
                    
                    set({ 
                        user: {
                            username: user.username,
                            email: user.email,
                            lastMoveDate: new Date(user.lastMoveDate)
                        } 
                    })
                    
                    return response.data
                } catch (error) {
                    console.error('Login failed:', error)
                    throw error
                }
            },
            
            register: async (registerRequest: RegisterRequest) => {
                try {
                    const response = await CAxios.post<RegisterResponse>('/auth/register', registerRequest)
                    const { token, user } = response.data
                    
                    set({ 
                        user: {
                            username: user.username,
                            email: user.email,
                            lastMoveDate: new Date(user.lastMoveDate)
                        } 
                    })
                    
                    return response.data
                } catch (error) {
                    console.error('Registration failed:', error)
                    throw error
                }
            },
            
            cooldownLeftInSeconds: async (cooldownRequest: CooldownLeftInSecondsRequest) => {
                try {
                    const response = await CAxios.post<CooldownLeftInSecondsResponse>(
                        '/user/cooldown', 
                        cooldownRequest
                    )
                    
                    const cooldownSeconds = response.data.cooldownLeftInSeconds
                    get().updateCooldown(cooldownSeconds)
                    
                    return response.data
                } catch (error) {
                    console.error('Failed to fetch cooldown:', error)
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
                const newCoolDate = Date.now() + (seconds * 1000)
                set({ coolDate: newCoolDate })
                
                if (get().user) {
                    set({ 
                        user: { 
                            ...get().user!,
                            lastMoveDate: new Date() 
                        } 
                    })
                }
            }
        }),
        {
            name: 'user-storage',
            partialize: (state) => ({ user: state.user }),
        }
    )
)