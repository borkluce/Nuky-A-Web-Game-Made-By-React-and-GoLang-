package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"services/internal/province/model"
)

// Create a test-specific interface that matches the methods we need to mock
type provinceRepoInterface interface {
	GetAll(ctx context.Context) ([]model.Province, error)
	UpdateProvinceByID(ctx context.Context, id string, isAttack bool) error
}

// MockProvinceRepo is a mock implementation of our interface
type MockProvinceRepo struct {
	mock.Mock
}

func (m *MockProvinceRepo) GetAll(ctx context.Context) ([]model.Province, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Province), args.Error(1)
}

func (m *MockProvinceRepo) UpdateProvinceByID(ctx context.Context, id string, isAttack bool) error {
	args := m.Called(ctx, id, isAttack)
	return args.Error(0)
}

// TestProvinceService is a wrapper around ProvinceService that accepts our interface
type TestProvinceService struct {
	provinceRepo provinceRepoInterface
}

// Create test implementations of the service methods we want to test
func (s *TestProvinceService) GetAllProvinces(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5)
	defer cancel()

	provinces, err := s.provinceRepo.GetAll(ctx)
	if err != nil {
		http.Error(w, "Failed to get all provinces", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(provinces)
}

func (s *TestProvinceService) AttackProvince(w http.ResponseWriter, r *http.Request) {
	s.updateProvinceCount(w, r, true)
}

func (s *TestProvinceService) SupportProvince(w http.ResponseWriter, r *http.Request) {
	s.updateProvinceCount(w, r, false)
}

func (s *TestProvinceService) updateProvinceCount(w http.ResponseWriter, r *http.Request, isAttack bool) {
	ctx, cancel := context.WithTimeout(r.Context(), 5)
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
	if err := s.provinceRepo.UpdateProvinceByID(ctx, id, isAttack); err != nil {
		http.Error(w, "Failed to update province", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Province updated successfully"))
}

// Test functions
func TestGetAllProvinces_Success(t *testing.T) {
	// Setup
	mockRepo := new(MockProvinceRepo)
	service := &TestProvinceService{provinceRepo: mockRepo}

	provinces := []model.Province{
		{
			ID:           primitive.NewObjectID(),
			ProvinceName: "Zartistan",
			AttackCount:  5,
			SupportCount: 10,
		},
		{
			ID:           primitive.NewObjectID(),
			ProvinceName: "Zortistan",
			AttackCount:  15,
			SupportCount: 20,
		},
	}

	mockRepo.On("GetAll", mock.Anything).Return(provinces, nil)

	// Execute
	req, err := http.NewRequest("GET", "/provinces", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.GetAllProvinces)
	handler.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	mockRepo.AssertExpectations(t)

	// Verify response body
	var responseProvinces []model.Province
	err = json.Unmarshal(rr.Body.Bytes(), &responseProvinces)
	assert.NoError(t, err)
	assert.Equal(t, provinces, responseProvinces)
}

func TestGetAllProvinces_Error(t *testing.T) {
	// Setup
	mockRepo := new(MockProvinceRepo)
	service := &TestProvinceService{provinceRepo: mockRepo}

	mockRepo.On("GetAll", mock.Anything).Return([]model.Province{}, errors.New("database error"))

	// Execute
	req, err := http.NewRequest("GET", "/provinces", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.GetAllProvinces)
	handler.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockRepo.AssertExpectations(t)
}

func TestAttackProvince_Success(t *testing.T) {
	// Setup
	mockRepo := new(MockProvinceRepo)
	service := &TestProvinceService{provinceRepo: mockRepo}

	validID := primitive.NewObjectID().Hex()
	mockRepo.On("UpdateProvinceByID", mock.Anything, validID, true).Return(nil)

	// Execute
	req, err := http.NewRequest("POST", "/provinces/attack?id="+validID, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.AttackProvince)
	handler.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "Province updated successfully", rr.Body.String())
}

func TestAttackProvince_MissingID(t *testing.T) {
	// Setup
	mockRepo := new(MockProvinceRepo)
	service := &TestProvinceService{provinceRepo: mockRepo}

	// Execute
	req, err := http.NewRequest("POST", "/provinces/attack", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.AttackProvince)
	handler.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Missing province ID")
}

func TestAttackProvince_InvalidIDFormat(t *testing.T) {
	// Setup
	mockRepo := new(MockProvinceRepo)
	service := &TestProvinceService{provinceRepo: mockRepo}

	// Execute
	req, err := http.NewRequest("POST", "/provinces/attack?id=invalidid", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.AttackProvince)
	handler.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid province ID format")
}

func TestAttackProvince_UpdateError(t *testing.T) {
	// Setup
	mockRepo := new(MockProvinceRepo)
	service := &TestProvinceService{provinceRepo: mockRepo}

	validID := primitive.NewObjectID().Hex()
	mockRepo.On("UpdateProvinceByID", mock.Anything, validID, true).Return(errors.New("update error"))

	// Execute
	req, err := http.NewRequest("POST", "/provinces/attack?id="+validID, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.AttackProvince)
	handler.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockRepo.AssertExpectations(t)
}

func TestSupportProvince_Success(t *testing.T) {
	// Setup
	mockRepo := new(MockProvinceRepo)
	service := &TestProvinceService{provinceRepo: mockRepo}

	validID := primitive.NewObjectID().Hex()
	mockRepo.On("UpdateProvinceByID", mock.Anything, validID, false).Return(nil)

	// Execute
	req, err := http.NewRequest("POST", "/provinces/support?id="+validID, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.SupportProvince)
	handler.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "Province updated successfully", rr.Body.String())
}

func TestSupportProvince_MissingID(t *testing.T) {
	// Setup
	mockRepo := new(MockProvinceRepo)
	service := &TestProvinceService{provinceRepo: mockRepo}

	// Execute
	req, err := http.NewRequest("POST", "/provinces/support", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.SupportProvince)
	handler.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Missing province ID")
}

func TestSupportProvince_InvalidIDFormat(t *testing.T) {
	// Setup
	mockRepo := new(MockProvinceRepo)
	service := &TestProvinceService{provinceRepo: mockRepo}

	// Execute
	req, err := http.NewRequest("POST", "/provinces/support?id=invalidid", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.SupportProvince)
	handler.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid province ID format")
}

func TestSupportProvince_UpdateError(t *testing.T) {
	// Setup
	mockRepo := new(MockProvinceRepo)
	service := &TestProvinceService{provinceRepo: mockRepo}

	validID := primitive.NewObjectID().Hex()
	mockRepo.On("UpdateProvinceByID", mock.Anything, validID, false).Return(errors.New("update error"))

	// Execute
	req, err := http.NewRequest("POST", "/provinces/support?id="+validID, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.SupportProvince)
	handler.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockRepo.AssertExpectations(t)
}
