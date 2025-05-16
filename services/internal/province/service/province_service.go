package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"services/internal/province/repo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProvinceService struct {
	repo *repo.ProvinceRepo
}

func NewProvinceService(repo *repo.ProvinceRepo) *ProvinceService {
	return &ProvinceService{repo: repo}
}

// GetAllProvinces returns all provinces as JSON
func (ps *ProvinceService) GetAllProvinces(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	provinces, err := ps.repo.GetAll(ctx)
	if err != nil {
		http.Error(w, "Failed to get all provinces", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(provinces)
}

// AttackProvince increases attack count by 1
func (ps *ProvinceService) AttackProvince(w http.ResponseWriter, r *http.Request) {
	ps.updateProvinceCount(w, r, true)
}

// SupportProvince increases support count by 1
func (ps *ProvinceService) SupportProvince(w http.ResponseWriter, r *http.Request) {
	ps.updateProvinceCount(w, r, false)
}

// shared logic for updating attack or support count
func (ps *ProvinceService) updateProvinceCount(w http.ResponseWriter, r *http.Request, isAttack bool) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Extract ID from query parameters
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing province ID", http.StatusBadRequest)
		return
	}

	// Validate ObjectID format
	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		http.Error(w, "Invalid province ID format", http.StatusBadRequest)
		return
	}

	// Call repo's UpdateProvince
	if err := ps.repo.UpdateProvinceByID(ctx, id, isAttack); err != nil {
		http.Error(w, "Failed to update province", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Province updated successfully"))
}
