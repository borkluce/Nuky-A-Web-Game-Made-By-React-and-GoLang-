export type LoginRequest = {
    username: string | null
    email: string | null
    password: string
}

export type LoginResponse = {
    token: string
    user: {
        username: string
        email: string
        lastMoveDate: string  // ISO date string from API
    }
}

// --------------------------------------------------------------------

export type RegisterRequest = {
    username: string
    email: string
    password: string
}

export type RegisterResponse = {
    token: string
    user: {
        username: string
        email: string
        lastMoveDate: string  // ISO date string from API
    }
}

// --------------------------------------------------------------------

export type CooldownLeftInSecondsRequest = {
    token: string
}

export type CooldownLeftInSecondsResponse = {
    cooldownLeftInSeconds: number
}

// --------------------------------------------------------------------

export type UserResponse = {
    username: string
    email: string
    lastMoveDate: string  // ISO date string from API
}