package services

import (
	"yab-explorer/models"
	"yab-explorer/repository"
)

type OrderService interface {
	GetOrder(orderID int) (models.Order, error)
	GetOrders(page, limit int) ([]models.Order, error)
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func New(repository repository.OrderRepository) OrderService {
	return &orderService{
		orderRepository: repository,
	}
}

func (s *orderService) GetOrder(orderID int) (models.Order, error) {
	order, err := s.orderRepository.GetOrder(orderID)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (s *orderService) GetOrders(page, limit int) ([]models.Order, error) {
	orders, err := s.orderRepository.GetOrders(page, limit)
	if err != nil {
		return []models.Order{}, err
	}
	return orders, nil
}
