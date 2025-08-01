package order

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type OrderHandler struct{}

func New() OrderHandler {
	return OrderHandler{}
}

func (h *OrderHandler) CalculatePacks(c echo.Context) error {

	return c.JSON(http.StatusOK, "ok")
}
