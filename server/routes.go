package server

// We add and load the routes
func (x *Server) API() {
	x.Echo.Static("/", "./public")

	api := x.Echo.Group("/api")
	api.POST("/pack-sizes", x.Handler.Order.SetPackSizes).Name = "post_pack_sizes"
	api.GET("/pack-sizes", x.Handler.Order.GetPackSizes).Name = "post_pack_sizes"
	api.GET("/calculate-packs", x.Handler.Order.CalculatePacks).Name = "get_calculate_packs"
}
