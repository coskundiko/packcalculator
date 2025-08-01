package server

// API routes
func (x *Server) API() {
	x.Echo.Static("/", "./public")

	api := x.Echo.Group("/api")
	api.GET("/pack-sizes", x.Handler.Order.GetPackSizes).Name = "post_pack_sizes"
	api.POST("/pack-sizes", x.Handler.Order.SetPackSizes).Name = "post_pack_sizes"
	api.POST("/calculate-packs", x.Handler.Order.CalculatePacks).Name = "post_calculate_packs"
}
