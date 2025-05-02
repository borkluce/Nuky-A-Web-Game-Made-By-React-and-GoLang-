export type LoginRequest = {
    username: string | null
    email: string | null
    password: string
}

export type LoginResponse = {
    token: string
}

// --------------------------------------------------------------------

export type RegisterRequest = {
    username: string
    email: string
    password: string
}

export type RegisterResponse = {
    token: string
}
