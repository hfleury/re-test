package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hfleury/re-test/internal/domain"
	"github.com/hfleury/re-test/internal/models"
)

type ConfigHandler struct {
	ConfigService domain.ConfigService
}

func NewConfigHandler(configService domain.ConfigService) *ConfigHandler {
	return &ConfigHandler{
		ConfigService: configService,
	}
}
func (h *ConfigHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/config/pack_sizes", h.UpdatePackSizes)
}

func (h *ConfigHandler) UpdatePackSizes(c *gin.Context) {
	var requestData models.UpdateConfigRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := h.ConfigService.UpdatePackSizes(requestData.PackSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pack sizes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pack sizes updated successfully"})
}
