package service

import (
	// Standart
	"encoding/json"
	"net/http"
	"time"

	// Internal
	"services/internal/auth/model"
	"services/internal/auth/repo"

	// Third
	"github.com/kahlery/pkg/go/auth/token"
	"github.com/kahlery/pkg/go/log/util"
)

type AuthService struct {
	userRepo *repo.UserRepo
}

func NewAuthService(userRepo *repo.UserRepo) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

// Services --------------------------------------------------------------------

// RegisterHandler handles user registration
func (as AuthService) Register(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.LogError("Failed to decode request body: "+err.Error(), "AuthService.RegisterHandler", "")
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Validate request
	if req.Username == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "Username, email, and password are required", http.StatusBadRequest)
		return
	}

	// Check if user already exists
	_, err := as.userRepo.GetUserByEmail(r.Context(), req.Email)
	if err == nil {
		http.Error(w, "Email already registered", http.StatusConflict)
		return
	}

	// Check if username is taken
	_, err = as.userRepo.GetUserByUsername(r.Context(), req.Username)
	if err == nil {
		http.Error(w, "Username already taken", http.StatusConflict)
		return
	}

	// Hash password
	hashedPassword, err := token.HashPassword(req.Password)
	if err != nil {
		util.LogError("Failed to hash password: "+err.Error(), "AuthService.RegisterHandler", "")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Create new user
	user := model.User{
		Username:     req.Username,
		Email:        req.Email,
		Password:     hashedPassword,
		LastMoveDate: time.Now(),
	}

	// Save user to database
	id, err := as.userRepo.CreateUser(r.Context(), user)
	if err != nil {
		util.LogError("Failed to create user: "+err.Error(), "AuthService.RegisterHandler", "")
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Set ID for response
	user.ID = id
	user.Password = "" // Don't send back password

	// Generate JWT token
	jwtToken, err := token.GenerateToken(user.ID.Hex())
	if err != nil {
		util.LogError("Failed to generate token: "+err.Error(), "AuthService.RegisterHandler", "")
		http.Error(w, "Failed to generate authentication token", http.StatusInternalServerError)
		return
	}

	// Send response
	response := model.RegisterResponse{
		Token: jwtToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// --------------------------------------------------------------------

// LoginHandler handles user login
func (as AuthService) Login(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.LogError("Failed to decode request body: "+err.Error(), "AuthService.LoginHandler", "")
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Validate request
	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Find user by email
	user, err := as.userRepo.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Verify password
	if err := token.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Update last move date
	user.LastMoveDate = time.Now()
	if err := as.userRepo.PutUser(r.Context(), *user); err != nil {
		util.LogError("Failed to update last move date: "+err.Error(), "AuthService.LoginHandler", "")
	}

	// Generate JWT token
	jwtToken, err := token.GenerateToken(user.ID.Hex())
	if err != nil {
		util.LogError("Failed to generate token: "+err.Error(), "AuthService.LoginHandler", "")
		http.Error(w, "Failed to generate authentication token", http.StatusInternalServerError)
		return
	}

	// Remove password from response
	user.Password = ""

	// Send response
	response := model.LoginResponse{
		Token: jwtToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
