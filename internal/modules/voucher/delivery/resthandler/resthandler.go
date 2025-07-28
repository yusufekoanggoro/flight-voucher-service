package resthandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yusufekoanggoro/flight-voucher-service/internal/factory/interfaces"
	"github.com/yusufekoanggoro/flight-voucher-service/internal/modules/voucher/domain/request"
	"github.com/yusufekoanggoro/flight-voucher-service/internal/modules/voucher/usecase"
	"github.com/yusufekoanggoro/flight-voucher-service/utils"
)

type RestHandler struct {
	uc usecase.VoucherUsecase
}

func NewRestHandler(uc usecase.VoucherUsecase) interfaces.RestHandler {
	return &RestHandler{
		uc: uc,
	}
}

func (h *RestHandler) RegisterRoutes(router gin.IRoutes) {
	router.GET("/check", h.CheckFlight)
	router.POST("/generate", h.GenerateVoucher)
}

func (h *RestHandler) CheckFlight(c *gin.Context) {
	var req request.CheckFlightRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request: "+err.Error())
		return
	}

	result, err := h.uc.CheckFlight(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exists": result,
	})
}

func (h *RestHandler) GenerateVoucher(c *gin.Context) {
	var req request.GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request: "+err.Error())
		return
	}

	result, err := h.uc.GenerateVoucher(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"seats":   result.Seats,
	})
}
