//go:build e2e
// +build e2e

// This file contains end-to-end tests for the PackSizeHandler.
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hfleury/re-test/internal/models"
	"github.com/hfleury/re-test/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePackSizes_E2E(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// real service with actual pack sizes
	service := services.NewPackSizeService([]int{23, 31, 53})
	handler := NewPackSizeHandler(service)

	router := gin.Default()
	handler.RegisterRoutes(router)

	req, _ := http.NewRequest("GET", "/calculate?order_amount=500000", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"message":"success"`)
	assert.Contains(t, w.Body.String(), `"packs"`)

	var response models.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	data, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)

	packsRaw, exists := data["packs"]
	assert.True(t, exists)

	packs := map[int]int{}
	for k, v := range packsRaw.(map[string]interface{}) {
		var keyInt int
		_, err := fmt.Sscanf(k, "%d", &keyInt)
		assert.NoError(t, err)

		packs[keyInt] = int(v.(float64))
	}

	expected := map[int]int{
		53: 9433,
		31: 1,
		23: 1,
	}
	assert.Equal(t, expected, packs)
}
