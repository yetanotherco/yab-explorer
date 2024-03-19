package services

import (
	"yab-explorer/models"
	"yab-explorer/repository"

	log "github.com/sirupsen/logrus"
)

type OrderService interface {
	GetOrder(orderID int) (models.Order, error)
	GetOrders(page, limit int) ([]models.Order, error)
}

type OrderServiceImpl struct {
	orderRepository repository.OrderRepository
}

func OrderServiceInit(orderRepository repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{orderRepository: orderRepository}
}

func (o OrderServiceImpl) GetOrder(orderID int) (models.Order, error) {
	log.Info("Called GetOrder with orderID: ", orderID, " in OrderServiceImpl.")
	order, err := o.orderRepository.GetOrder(orderID)
	if err != nil {
		log.Error("Error getting order with orderID: ", orderID, " in OrderServiceImpl. Error: ", err)
		return models.Order{}, err
	}
	return order, nil
}

func (o OrderServiceImpl) GetOrders(page, limit int) ([]models.Order, error) {
	log.Info("Called GetOrders with page: ", page, " and limit: ", limit, " in OrderServiceImpl.")
	orders, err := o.orderRepository.GetOrders(page, limit)
	if err != nil {
		log.Error("Error getting orders with page: ", page, " and limit: ", limit, " in OrderServiceImpl. Error: ", err)
		return []models.Order{}, err
	}
	return orders, nil
}
