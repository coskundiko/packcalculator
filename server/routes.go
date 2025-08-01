package server

// We add and load the routes
func (x *Server) API() {
	x.Echo.Static("/store", "store").Name = "static_store"
	x.Echo.Static("/", "./public")
	//x.Echo.GET("/", func(c echo.Context) error {
	//	return c.JSON(http.StatusOK, "OK - Mode: ")
	//})

	x.Echo.GET("/calculate-packs", x.Handler.Order.CalculatePacks).Name = "get_calculate_packs"

	//api := e.Group("/api")
}
