package services

import (
	"math"
	"yab-explorer/domain/dtos"
	"yab-explorer/domain/models"
	"yab-explorer/repository"

	log "github.com/sirupsen/logrus"
)

type OrderService interface {
	GetOrder(orderId int) (models.Order, error)
	GetOrders(page, pageSize int, sort, direction string) (dtos.PaginatedSearchResultDto, error)
	GetOrdersCount() (int, error)
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

func (o OrderServiceImpl) GetOrders(page, pageSize int, sort, direction string) (dtos.PaginatedSearchResultDto, error) {
	log.Info("Called GetOrders with page: ", page, " and pageSize: ", pageSize, " with sort: ", sort, " and direction: ", direction, " in OrderServiceImpl.")

	orders, err := o.orderRepository.GetOrders(page, pageSize, sort, direction)
	orderCount, _ := o.orderRepository.GetTotalOrders()
	totalPages := int(math.Ceil(float64(orderCount) / float64(pageSize)))

	if orderCount == 0 || page > totalPages {
		return *dtos.NewPaginatedSearchResultDto(0, pageSize, []models.Order{}, 0), nil
	}

	if err != nil {
		log.Error("Error getting orders with page: ", page, " and pageSize: ", pageSize, " with sort: ", sort, " and direction: ", direction, " in OrderServiceImpl. Error: ", err)
		return dtos.PaginatedSearchResultDto{}, err
	}

	var ordersDto []dtos.OrderDto

	for _, order := range orders {
		ordersDto = append(ordersDto, dtos.OrderToDto(order))
	}

	return *dtos.NewPaginatedSearchResultDto(page, pageSize, ordersDto, orderCount), nil
}

func (o OrderServiceImpl) GetOrdersCount() (int, error) {
	log.Info("Called GetOrdersCount in OrderServiceImpl.")
	orderCount, err := o.orderRepository.GetTotalOrders()
	if err != nil {
		log.Error("Error getting orders count in OrderServiceImpl. Error: ", err)
		return 0, err
	}
	return orderCount, nil
}
