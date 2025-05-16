export type User = {
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
    userData: { username: string; email: string }
): User => ({
    username: userData.username,
    email: userData.email,
    last_move_date: new Date(),
    token,
    isAuthenticated: true,
})
