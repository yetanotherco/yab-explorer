package configs

import (
	"yab-explorer/controllers"
	"yab-explorer/repository"
	"yab-explorer/services"
)

type Initialization struct {
	orderRepository repository.OrderRepository
	orderService    services.OrderService
	OrderController controllers.OrderController
}

func NewInitialization(orderRepository repository.OrderRepository,
	orderService services.OrderService,
	OrderController controllers.OrderController) *Initialization {
	return &Initialization{
		orderRepository: orderRepository,
		orderService:    orderService,
		OrderController: OrderController,
	}
}
