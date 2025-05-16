package model

// RegisterRequest represents the data needed for user registration.
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

// --------------------------------------------------------------------

// LoginRequest represents the data needed for user login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// --------------------------------------------------------------------

type CooldownLeftInSecondsRequest struct {
	Token string `json:"token"`
}

type CooldownLeftInSecondsResponse struct {
	CoolDownLeftInSeconds int32 `json:"cooldown_left_in_seconds"`
}
