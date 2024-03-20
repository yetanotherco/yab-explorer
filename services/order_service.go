package services

import (
	"yab-explorer/models"
	"yab-explorer/repository"

	log "github.com/sirupsen/logrus"
)

type OrderService interface {
	GetOrder(orderID int) (models.Order, error)
	GetOrders(page, pageSize int) ([]models.Order, error)
	GetTotalOrders() (int, error)
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

func (o OrderServiceImpl) GetOrders(page, pageSize int) ([]models.Order, error) {
	log.Info("Called GetOrders with page: ", page, " and pageSize: ", pageSize, " in OrderServiceImpl.")
	orders, err := o.orderRepository.GetOrders(page, pageSize)
	if err != nil {
		log.Error("Error getting orders with page: ", page, " and pageSize: ", pageSize, " in OrderServiceImpl. Error: ", err)
		return []models.Order{}, err
	}
	return orders, nil
}

func (o OrderServiceImpl) GetTotalOrders() (int, error) {
	log.Info("Called GetTotalOrders in OrderServiceImpl.")
	totalOrders, err := o.orderRepository.GetTotalOrders()
	if err != nil {
		log.Error("Error getting total orders in OrderServiceImpl. Error: ", err)
		return 0, err
	}
	return totalOrders, nil
}
