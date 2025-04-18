package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"services/internal/province/repo"
)

type ProvinceService struct {
	// Fields/Properties --------------------------------------------------------------------
	repo *repo.ProvinceRepo
}

// Constructor --------------------------------------------------------------------
func NewProvinceService(repo *repo.ProvinceRepo) *ProvinceService {
	return &ProvinceService{repo: repo}
}

// Behaviours --------------------------------------------------------------------
func (ph *ProvinceService) GetAllProvinces(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	provinces, err := ph.repo.GetAll(ctx)
	if err != nil {
		http.Error(w, "Failed to get all provinces", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(provinces)
}

// func ()  AttackProvince
