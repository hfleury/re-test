package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hfleury/re-test/internal/domain"
	"github.com/hfleury/re-test/internal/models"
)

type PackSizeHandler struct {
	packService   domain.PackSizeService
	configService domain.ConfigService
}

func NewPackSizeHandler(packService domain.PackSizeService, configService domain.ConfigService) *PackSizeHandler {
	return &PackSizeHandler{
		packService:   packService,
		configService: configService,
	}
}

func (h *PackSizeHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/calculate", h.CalculatePackSizes)
}

func (h *PackSizeHandler) CalculatePackSizes(c *gin.Context) {
	orderItems := c.Query("order_amount")
	if orderItems == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "order_amount is required",
		})
		return
	}

	amount, err := strconv.Atoi(orderItems)
	if err != nil || amount <= 0 {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "order_amount must be a positive integer",
		})
		return
	}

	packSizes := h.configService.GetPackSizes()

	result, err := h.packService.CalculatePackSizeByOrderAmount(amount, packSizes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data: models.CalculateResponse{
			Packs: result,
		},
	})
}
