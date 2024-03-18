package repository

import (
	"yab-explorer/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrder(orderID int) (models.Order, error)
	GetOrders(page, limit int) ([]models.Order, error)
}

type orderRepositry struct {
	connection *gorm.DB
}

func NewOrderRepository(connection *gorm.DB) OrderRepository {
	return &orderRepositry{connection: connection}
}

func (or *orderRepositry) GetOrder(orderID int) (models.Order, error) {
	var order models.Order
	err := or.connection.First(&order, orderID).Error
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (or *orderRepositry) GetOrders(page, limit int) ([]models.Order, error) {
	var orders []models.Order
	err := or.connection.Limit(limit).Offset((page - 1) * limit).Order("created_at desc").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
