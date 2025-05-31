export type User = {
    id?: string
    username: string
    email: string
    last_move_date: Date

    // Authentication related
    token?: string
    isAuthenticated?: boolean

    // Game related
    cooldownLeftInSeconds?: number
}

// Helper functions
export const createEmptyUser = (): User => ({
    username: "",
    email: "",
    last_move_date: new Date(),
    isAuthenticated: false,
})

export const createUserFromToken = (
    token: string,
    userData: { 
        id?: string
        username: string
        email: string
        last_move_date?: Date
    }
): User => ({
    id: userData.id,
    username: userData.username,
    email: userData.email,
    last_move_date: userData.last_move_date || new Date(),
    token,
    isAuthenticated: true,
})

export const createUserFromResponse = (
    token: string,
    user: {
        ID?: string
        username: string
        email: string
        last_move_date?: string | Date
    }
): User => ({
    id: user.ID,
    username: user.username,
    email: user.email,
    last_move_date: user.last_move_date 
        ? (typeof user.last_move_date === 'string' 
            ? new Date(user.last_move_date) 
            : user.last_move_date)
        : new Date(),
    token,
    isAuthenticated: true,
})