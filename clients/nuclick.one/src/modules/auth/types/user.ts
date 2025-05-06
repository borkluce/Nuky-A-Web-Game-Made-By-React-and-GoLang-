export type User = {
    username: string
    email: string
    lastMoveDate: Date
    
    // Authentication related
    token?: string
    isAuthenticated?: boolean
    
    // Game related 
    cooldownLeftInSeconds?: number
}

// Helper functions
export const createEmptyUser = (): User => ({
    username: '',
    email: '',
    lastMoveDate: new Date(),
    isAuthenticated: false
})

export const createUserFromToken = (token: string, userData: { username: string, email: string }): User => ({
    username: userData.username,
    email: userData.email,
    lastMoveDate: new Date(),
    token,
    isAuthenticated: true
})