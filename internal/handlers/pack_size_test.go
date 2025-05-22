package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Adjust this to accept both mocks, always pass both
func setupRouterWithHandler(packService *MockPackSizeService, configService *MockConfigService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := NewPackSizeHandler(packService, configService)
	handler.RegisterRoutes(router)
	return router
}

func TestCalculatePackSizes_MissingQueryParam(t *testing.T) {
	mockPackService := new(MockPackSizeService)
	mockConfigService := new(MockConfigService)

	router := setupRouterWithHandler(mockPackService, mockConfigService)

	req, _ := http.NewRequest("GET", "/calculate", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "order_amount is required")
}

func TestCalculatePackSizes_InvalidQueryParam(t *testing.T) {
	mockPackService := new(MockPackSizeService)
	mockConfigService := new(MockConfigService)

	router := setupRouterWithHandler(mockPackService, mockConfigService)

	req, _ := http.NewRequest("GET", "/calculate?order_amount=abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "order_amount must be a positive integer")
}

func TestCalculatePackSizes_ServiceError(t *testing.T) {
	mockConfigService := new(MockConfigService)
	mockPackSizeService := new(MockPackSizeService)

	mockConfigService.On("GetPackSizes").Return([]int{250, 500})
	mockPackSizeService.On("CalculatePackSizeByOrderAmount", 100, []int{250, 500}).Return(map[int]int{}, errors.New("some internal error"))

	router := setupRouterWithHandler(mockPackSizeService, mockConfigService)

	req, _ := http.NewRequest("GET", "/calculate?order_amount=100", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "internal server error")

	mockConfigService.AssertExpectations(t)
	mockPackSizeService.AssertExpectations(t)
}

func TestCalculatePackSizes_Success(t *testing.T) {
	mockConfigService := new(MockConfigService)
	mockPackSizeService := new(MockPackSizeService)

	mockConfigService.On("GetPackSizes").Return([]int{250, 500})

	expected := map[int]int{
		250: 2,
		500: 1,
	}

	mockPackSizeService.On("CalculatePackSizeByOrderAmount", 1000, []int{250, 500}).Return(expected, nil)

	router := setupRouterWithHandler(mockPackSizeService, mockConfigService)

	req, _ := http.NewRequest("GET", "/calculate?order_amount=1000", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"message":"success"`)
	assert.Contains(t, w.Body.String(), `"packs":{"250":2,"500":1}`)

	mockConfigService.AssertExpectations(t)
	mockPackSizeService.AssertExpectations(t)
}
