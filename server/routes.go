package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (x *Server) API() {
	x.Echo.Static("/store", "store").Name = "static_store"
	x.Echo.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK - Mode: ")
	})
	//api := e.Group("/api")
	x.Echo.Static("/diko", "./server/assets")
}
