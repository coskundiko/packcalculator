package server

// We add and load the routes
func (x *Server) API() {
	x.Echo.Static("/", "./public")

	api := x.Echo.Group("/api")
	api.POST("/calculate-packs", x.Handler.Order.CalculatePacks).Name = "get_calculate_packs"
}
