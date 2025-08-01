package order

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"packcalculator/pkg/calculator"
)

type OrderHandler struct{}

type CalculatePacksRequest struct {
	OrderSize int   `json:"order_size"`
	PackSizes []int `json:"pack_sizes"`
}

func New() OrderHandler {
	return OrderHandler{}
}

func (h *OrderHandler) CalculatePacks(c echo.Context) error {
	req := CalculatePacksRequest{}

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	} else if req.OrderSize <= 0 {
		return c.JSON(http.StatusBadRequest, "order_size must be a positive number")
	} else if len(req.PackSizes) == 0 {
		return c.JSON(http.StatusBadRequest, "order_size must be a positive number")
	}

	packs := calculator.CalculatePacks(req.OrderSize, req.PackSizes)

	return c.JSON(http.StatusOK, packs)
}
