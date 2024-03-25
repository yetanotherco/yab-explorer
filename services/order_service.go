package services

import (
	"math"
	"yab-explorer/models"
	"yab-explorer/repository"

	log "github.com/sirupsen/logrus"
)

type OrderService interface {
	GetOrder(orderId int) (models.Order, error)
	GetOrders(page, pageSize int, sort, direction string) (models.PaginatedSearchResult, error)
}

type OrderServiceImpl struct {
	orderRepository repository.OrderRepository
}

func OrderServiceInit(orderRepository repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{orderRepository: orderRepository}
}

func (o OrderServiceImpl) GetOrder(orderId int) (models.Order, error) {
	log.Info("Called GetOrder with orderId: ", orderId, " in OrderServiceImpl.")
	order, err := o.orderRepository.GetOrder(orderId)
	if err != nil {
		log.Error("Error getting order with orderId: ", orderId, " in OrderServiceImpl. Error: ", err)
		return models.Order{}, err
	}
	return order, nil
}

func (o OrderServiceImpl) GetOrders(page, pageSize int, sort, direction string) (models.PaginatedSearchResult, error) {
	log.Info("Called GetOrders with page: ", page, " and pageSize: ", pageSize, " with sort: ", sort, " and direction: ", direction, " in OrderServiceImpl.")

	orders, err := o.orderRepository.GetOrders(page, pageSize, sort, direction)
	orderCount, _ := o.orderRepository.GetTotalOrders()
	totalPages := int(math.Ceil(float64(orderCount) / float64(pageSize)))

	if orderCount == 0 || page > totalPages {
		return *models.NewPaginatedSearchResult(0, pageSize, []models.Order{}, 0), nil
	}

	if err != nil {
		log.Error("Error getting orders with page: ", page, " and pageSize: ", pageSize, " with sort: ", sort, " and direction: ", direction, " in OrderServiceImpl. Error: ", err)
		return models.PaginatedSearchResult{}, err
	}

	return *models.NewPaginatedSearchResult(page, pageSize, orders, orderCount), nil
}
