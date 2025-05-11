package service

import (
	"encoding/json"
	"net/http"
	"services/internal/auth/model"
	"services/internal/auth/repo"
	"time"

	"github.com/kahleryasla/pkg/go/log/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repo.UserRepo
}

func NewAuthService(userRepo *repo.UserRepo) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

// RegisterRequest represents the data needed for user registration
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest represents the data needed for user login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse represents the response for auth operations
type AuthResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Token   string      `json:"token,omitempty"`
	User    *model.User `json:"user,omitempty"`
}

// RegisterHandler handles user registration
func (as *AuthService) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req RegisterRequest
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		util.LogError("Failed to hash password: "+err.Error(), "AuthService.RegisterHandler", "")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Create new user
	user := model.User{
		Username:     req.Username,
		Email:        req.Email,
		Password:     string(hashedPassword),
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

	// Generate token
	token, err := generateToken(user.ID.Hex())
	if err != nil {
		util.LogError("Failed to generate token: "+err.Error(), "AuthService.RegisterHandler", "")
		http.Error(w, "Failed to generate authentication token", http.StatusInternalServerError)
		return
	}

	// Send response
	response := AuthResponse{
		Success: true,
		Message: "User registered successfully",
		Token:   token,
		User:    &user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// LoginHandler handles user login
func (as *AuthService) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req LoginRequest
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
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Update last move date
	user.LastMoveDate = time.Now()
	if err := as.userRepo.PutUser(r.Context(), *user); err != nil {
		util.LogError("Failed to update last move date: "+err.Error(), "AuthService.LoginHandler", "")
	}

	// Generate token
	token, err := generateToken(user.ID.Hex())
	if err != nil {
		util.LogError("Failed to generate token: "+err.Error(), "AuthService.LoginHandler", "")
		http.Error(w, "Failed to generate authentication token", http.StatusInternalServerError)
		return
	}

	// Remove password from response
	user.Password = ""

	// Send response
	response := AuthResponse{
		Success: true,
		Message: "Login successful",
		Token:   token,
		User:    user,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// VerifyHandler handles token verification
func (as *AuthService) VerifyHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get token from Authorization header
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization token required", http.StatusUnauthorized)
		return
	}

	// Remove "Bearer " prefix if present
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	// Verify token
	userID, err := verifyToken(token)
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Get user from database
	user, err := as.userRepo.GetUserByID(r.Context(), objectID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Remove password from response
	user.Password = ""

	// Send response
	response := AuthResponse{
		Success: true,
		Message: "Token verified",
		User:    user,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Helper functions for token generation and verification
func generateToken(userID string) (string, error) {
	// TODO: Implement proper JWT token generation here
	// For now, we'll return a simple token for demonstration
	return "token_" + userID, nil
}

func verifyToken(token string) (string, error) {
	// TODO: Implement proper JWT token verification here
	// For now, we'll extract the userID from the token for demonstration
	if len(token) <= 6 || token[:6] != "token_" {
		return "", http.ErrNoCookie // Using ErrNoCookie as a placeholder for invalid token error
	}
	return token[6:], nil
}
