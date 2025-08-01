package handler

import "packcalculator/handler/order"

type Handler struct {
	Order order.OrderHandler
}

func New() Handler {
	return Handler{
		Order: order.New(),
	}
}
