package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hfleury/re-test/internal/domain"
	"github.com/stretchr/testify/assert"
)

func setupConfigRouter(configService domain.ConfigService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := NewConfigHandler(configService)
	handler.RegisterRoutes(router)
	return router
}

func TestUpdatePackSizes_Success(t *testing.T) {
	mockConfigService := new(MockConfigService)

	mockConfigService.On("UpdatePackSizes", []int{100, 200, 300}).Return(nil)

	router := setupConfigRouter(mockConfigService)

	body := `{"pack_sizes": [100, 200, 300]}`
	req, _ := http.NewRequest(http.MethodPost, "/config/pack_sizes", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Pack sizes updated successfully"}`, w.Body.String())
	mockConfigService.AssertExpectations(t)
}

func TestUpdatePackSizes_InvalidJSON(t *testing.T) {
	mockConfigService := new(MockConfigService)

	router := setupConfigRouter(mockConfigService)

	body := `{"pack_sizes": "not_an_array"}`
	req, _ := http.NewRequest(http.MethodPost, "/config/pack_sizes", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "Invalid request payload"}`, w.Body.String())
}

func TestUpdatePackSizes_ServiceError(t *testing.T) {
	mockConfigService := new(MockConfigService)

	mockConfigService.On("UpdatePackSizes", []int{100, 200}).Return(errors.New("write error"))

	router := setupConfigRouter(mockConfigService)

	body := `{"pack_sizes": [100, 200]}`
	req, _ := http.NewRequest(http.MethodPost, "/config/pack_sizes", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"error": "Failed to update pack sizes"}`, w.Body.String())
	mockConfigService.AssertExpectations(t)
}
