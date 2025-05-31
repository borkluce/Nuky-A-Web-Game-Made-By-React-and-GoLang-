package model

// Request DTOs
type LoginRequest struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password string  `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CooldownLeftInSecondsRequest struct {
	Token string `json:"token"`
}

// Response DTOs
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user,omitempty"` // Optional: include user data
}

type RegisterResponse struct {
	Token string `json:"token"`
	User  User   `json:"user,omitempty"` // Optional: include user data
}

type CooldownLeftInSecondsResponse struct {
	CooldownLeftInSeconds int `json:"cooldown_left_in_seconds"`
}
