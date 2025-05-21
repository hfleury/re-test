package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hfleury/re-test/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockPackSizeService struct {
	mock.Mock
}

func (m *mockPackSizeService) CalculatePackSizeByOrderAmount(orderItems int) (map[int]int, error) {
	args := m.Called(orderItems)
	return args.Get(0).(map[int]int), args.Error(1)
}

func setupRouterWithHandler(service domain.PackSizeService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := NewPackSizeHandler(service)
	handler.RegisterRoutes(router)
	return router
}

func TestCalculatePackSizes_MissingQueryParam(t *testing.T) {
	service := new(mockPackSizeService)
	router := setupRouterWithHandler(service)

	req, _ := http.NewRequest("GET", "/calculate", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "order_amount is required")
}

func TestCalculatePackSizes_InvalidQueryParam(t *testing.T) {
	service := new(mockPackSizeService)
	router := setupRouterWithHandler(service)

	req, _ := http.NewRequest("GET", "/calculate?order_amount=abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "order_amount must be a positive integer")
}

func TestCalculatePackSizes_ServiceError(t *testing.T) {
	service := new(mockPackSizeService)
	service.On("CalculatePackSizeByOrderAmount", 100).Return(map[int]int{}, errors.New("some internal error"))

	router := setupRouterWithHandler(service)

	req, _ := http.NewRequest("GET", "/calculate?order_amount=100", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "internal server error")
}

func TestCalculatePackSizes_Success(t *testing.T) {
	service := new(mockPackSizeService)
	expected := map[int]int{
		250: 2,
		500: 1,
	}
	service.On("CalculatePackSizeByOrderAmount", 1000).Return(expected, nil)

	router := setupRouterWithHandler(service)

	req, _ := http.NewRequest("GET", "/calculate?order_amount=1000", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"message":"success"`)
	assert.Contains(t, w.Body.String(), `"packs":{"250":2,"500":1}`)
}
