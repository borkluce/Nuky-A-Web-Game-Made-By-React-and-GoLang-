package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"services/internal/province/model"
	"services/internal/province/repo"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kahlery/pkg/go/auth/token"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProvinceService struct {
	repo      *repo.ProvinceRepo
	startDate time.Time // Game start date
}

func NewProvinceService(repo *repo.ProvinceRepo, startDate time.Time) *ProvinceService {
	return &ProvinceService{
		repo:      repo,
		startDate: startDate,
	}
}

// --------------------------------------------------------------------
func (ps *ProvinceService) GetAllProvinces(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	provinces, err := ps.repo.GetAll(ctx)
	if err != nil {
		http.Error(w, "Failed to get all provinces", http.StatusInternalServerError)
		return
	}

	// Wrap the provinces in the DTO
	response := model.GetAllProvinceResponse{
		ProvinceList: provinces,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// GetTopProvinces returns the top 5 provinces by score difference (attackCount - supportCount)
func (ps *ProvinceService) GetTopProvinces(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Get provinces sorted by score difference
	provinces, err := ps.repo.GetProvincesByScoreDifference(ctx)
	if err != nil {
		http.Error(w, "Failed to get top provinces", http.StatusInternalServerError)
		return
	}

	// Take only the top 5 (or less if there are fewer than 5 provinces)
	topCount := 5
	if len(provinces) < topCount {
		topCount = len(provinces)
	}
	topProvinces := provinces[:topCount]

	// Wrap in response object
	response := struct {
		Provinces []model.Province `json:"provinces"`
	}{
		Provinces: topProvinces,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// AttackProvince increases attack count by 1
func (ps *ProvinceService) AttackProvince(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// This is for making sure if a logged in user is attacking
	/*	_, err := ExtractUserIDFromRequest(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		} */

	// Parse request body
	var req model.AttackProvinceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate province ID
	if req.ProvinceID == "" {
		http.Error(w, "Province ID is required", http.StatusBadRequest)
		return
	}

	// Validate ObjectID format
	if _, err := primitive.ObjectIDFromHex(req.ProvinceID); err != nil {
		http.Error(w, "Invalid province ID format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Update province attack count
	if err := ps.repo.UpdateProvinceByID(ctx, req.ProvinceID, true); err != nil {
		http.Error(w, "Failed to update province", http.StatusInternalServerError)
		return
	}

	response := model.AttackProvinceResponse{
		IsSuccess: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// --------------------------------------------------------------------
// SupportProvince increases support count by 1
func (ps *ProvinceService) SupportProvince(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// This is for making sure if a logged in user is supporting
	// Extract user ID from JWT token
	/*	_, err := ExtractUserIDFromRequest(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		} */

	// Parse request body
	var req model.SupportProvinceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate province ID
	if req.ProvinceID == "" {
		http.Error(w, "Province ID is required", http.StatusBadRequest)
		return
	}

	// Validate ObjectID format
	if _, err := primitive.ObjectIDFromHex(req.ProvinceID); err != nil {
		http.Error(w, "Invalid province ID format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Update province support count
	if err := ps.repo.UpdateProvinceByID(ctx, req.ProvinceID, false); err != nil {
		http.Error(w, "Failed to update province", http.StatusInternalServerError)
		return
	}

	response := model.SupportProvinceResponse{
		IsSuccess: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// --------------------------------------------------------------------
// UpdateDestroymentRound handles the nuke operation
func (ps *ProvinceService) UpdateDestroymentRound(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	// Calculate round count (days passed since start date)
	currentTime := time.Now()
	daysPassed := int(currentTime.Sub(ps.startDate).Hours() / 24)
	roundCount := daysPassed

	// Update destroyment round of the worst province (highest attackCount - supportCount)
	err := ps.repo.UpdateDestroymentRoundOfTheWorstProvince(ctx, roundCount)
	if err != nil {
		http.Error(w, "Failed to update destroyment round", http.StatusInternalServerError)
		return
	}

	// Reset all provinces' attack and support counts
	err = ps.repo.ResetAllProvinceCounts(ctx)
	if err != nil {
		http.Error(w, "Failed to reset province counts", http.StatusInternalServerError)
		return
	}

	response := struct {
		Message    string `json:"message"`
		RoundCount int    `json:"round_count"`
	}{
		Message:    "Destroyment round updated and counts reset successfully",
		RoundCount: roundCount,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// --------------------------------------------------------------------
// ExecuteDestroymentRound performs the nuke operation without HTTP context (for cron jobs)
func (ps *ProvinceService) ExecuteDestroymentRound(ctx context.Context) (int, error) {
	// Calculate round count (days passed since start date)
	currentTime := time.Now()
	daysPassed := int(currentTime.Sub(ps.startDate).Hours() / 24)
	roundCount := daysPassed

	// Update destroyment round of the worst province (highest attackCount - supportCount)
	err := ps.repo.UpdateDestroymentRoundOfTheWorstProvince(ctx, roundCount)
	if err != nil {
		return 0, err
	}

	// Reset all provinces' attack and support counts
	err = ps.repo.ResetAllProvinceCounts(ctx)
	if err != nil {
		return 0, err
	}

	return roundCount, nil
}

// --------------------------------------------------------------------
// GetCurrentRound returns wihch round the game is in
func (ps *ProvinceService) GetCurrentRound(ctx context.Context) (int, error) {
	// Calculate current round based on start date
	now := time.Now().UTC()
	daysSinceStart := int(now.Sub(ps.startDate).Hours() / 24)

	// If it's past 14:00 UTC today, we're in the next round
	todayAt14 := time.Date(now.Year(), now.Month(), now.Day(), 14, 0, 0, 0, time.UTC)
	if now.After(todayAt14) {
		daysSinceStart++
	}

	return daysSinceStart + 1, nil // Round starts from 1
}

func (ps *ProvinceService) GetCurrentRoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	roundCount, err := ps.GetCurrentRound(ctx)
	if err != nil {
		http.Error(w, "Failed to get current round", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := fmt.Sprintf(`{"round": %d, "success": true}`, roundCount)
	w.Write([]byte(response))
}

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
