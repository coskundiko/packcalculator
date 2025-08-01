package order

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

type OrderHandler struct {
	PackSizes []int `json:"pack_sizes"`
	sync.Mutex
}

type SetPackSizesRequest struct {
	PackSizes []int `json:"pack_sizes"`
}

type CalculatePacksRequest struct {
	OrderSize int `json:"order_size"`
}

func New() OrderHandler {
	return OrderHandler{
		PackSizes: []int{250, 500, 1000, 2000, 5000},
	}
}

func (h *OrderHandler) CalculatePacks(c echo.Context) error {
	req := CalculatePacksRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	if req.OrderSize <= 0 {
		return c.JSON(http.StatusBadRequest, "order_size must be a positive number")
	}

	//packs := calculator.CalculatePacks(req.OrderSize, req.PackSizes)

	return c.JSON(http.StatusOK, h.PackSizes)
}

func (h *OrderHandler) GetPackSizes(c echo.Context) error {
	return c.JSON(http.StatusOK, h.PackSizes)
}

func (h *OrderHandler) SetPackSizes(c echo.Context) error {
	req := new(SetPackSizesRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	if len(req.PackSizes) == 0 {
		return c.JSON(http.StatusBadRequest, "pack_sizes must be a positive number")
	}

	h.Mutex.Lock()
	h.PackSizes = req.PackSizes
	h.Mutex.Unlock()

	return c.JSON(http.StatusOK, h.PackSizes)
}
