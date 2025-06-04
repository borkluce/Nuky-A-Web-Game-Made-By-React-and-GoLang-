package service

import (
	// Standart
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	// Internal
	"services/internal/auth/model"
	"services/internal/auth/repo"

	// Third
	"github.com/golang-jwt/jwt/v4"
	"github.com/kahlery/pkg/go/auth/token"
	"github.com/kahlery/pkg/go/log/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// Check if user already exists by email
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
		User:  user,
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
	if (req.Email == nil && req.Username == nil) || req.Password == "" {
		http.Error(w, "Email/username and password are required", http.StatusBadRequest)
		return
	}

	var user *model.User
	var err error

	// Try to find user by email first, then by username
	if req.Email != nil && *req.Email != "" {
		user, err = as.userRepo.GetUserByEmail(r.Context(), *req.Email)
	} else if req.Username != nil && *req.Username != "" {
		user, err = as.userRepo.GetUserByUsername(r.Context(), *req.Username)
	}

	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Verify password
	if err := token.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
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
		User:  *user,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// --------------------------------------------------------------------

// POST /api/user/update-move-date
func (as AuthService) UpdateMoveDate(w http.ResponseWriter, r *http.Request) {
	userID, err := ExtractUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	objID, _ := primitive.ObjectIDFromHex(userID)
	user, err := as.userRepo.GetUserByID(r.Context(), objID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user.LastMoveDate = time.Now()
	err = as.userRepo.PutUser(r.Context(), *user)
	if err != nil {
		http.Error(w, "Failed to update", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// --------------------------------------------------------------------
// --------------------------------------------------------------------
// --------------------------------------------------------------------

// This part may be better if we import ExtractUserIDFromRequest into kahlery package

// ExtractUserIDFromRequest extracts user ID from the Authorization header
func ExtractUserIDFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return "", errors.New("invalid authorization header format")
	}

	tokenString := strings.TrimPrefix(authHeader, bearerPrefix)
	tokenString = strings.TrimSpace(tokenString)
	if tokenString == "" {
		return "", errors.New("token is empty")
	}

	// Use the token package's VerifyToken function instead
	jwtToken, err := token.VerifyToken(tokenString)
	if err != nil {
		return "", err
	}

	if !jwtToken.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	userID, ok := claims["userName"].(string)
	if !ok || userID == "" {
		return "", errors.New("userName not found in token")
	}

	return userID, nil
}
