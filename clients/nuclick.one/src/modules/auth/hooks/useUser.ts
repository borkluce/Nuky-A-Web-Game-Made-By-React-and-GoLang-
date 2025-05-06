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
    cooldown: number          // Seconds remaining in cooldown
    cooldownTotal: number     // Total cooldown period
    register: (registerRequest: RegisterRequest) => Promise<void>
    isAllowedToMove: () => boolean
    setCooldown: (seconds: number) => void
    resetCooldown: () => void
    startCooldownTimer: () => void
}

export const useUser = create<useUserState>((set, get) => ({
    user: null,
    cooldown: 0,          // Current cooldown seconds remaining
    cooldownTotal: 30,    // Default cooldown period is 30 seconds
    
    register: async (registerRequest) => {
        const response = await CAxios.post<RegisterRequest>(
            "/auth/register",
            registerRequest
        )
    },
    
    isAllowedToMove: () => {
        // Check if user can perform actions based on cooldown
        return get().cooldown <= 0
    },
    
    setCooldown: (seconds) => {
        set({ 
            cooldown: seconds,
            cooldownTotal: seconds 
        })
    },
    
    resetCooldown: () => {
        set({ cooldown: 0 })
    },
    
    startCooldownTimer: () => {
        // Set cooldown to default value
        const cooldownValue = get().cooldownTotal || 30
        set({ cooldown: cooldownValue })
        
        // Start countdown timer
        const intervalId = setInterval(() => {
            const currentCooldown = get().cooldown
            
            if (currentCooldown <= 1) {
                // Clear interval when cooldown reaches 0
                clearInterval(intervalId)
                set({ cooldown: 0 })
            } else {
                // Decrement cooldown
                set({ cooldown: currentCooldown - 1 })
            }
        }, 1000)
        
        // Return the interval ID in case we want to clear it elsewhere
        return intervalId
    }
}))