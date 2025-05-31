export type LoginRequest = {
    username: string | null
    email: string | null
    password: string
}

export type LoginResponse = {
    token: string
    user?: {
        ID: string
        username: string
        email: string
        last_move_date: string
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
    user?: {
        ID: string
        username: string
        email: string
        last_move_date: string
    }
}

// --------------------------------------------------------------------

export type CooldownLeftInSecondsRequest = {
    token: string
}

export type CooldownLeftInSecondsResponse = {
    cooldown_left_in_seconds: number
}