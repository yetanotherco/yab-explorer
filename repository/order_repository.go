package repository

import (
	"yab-explorer/models"

	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

type OrderRepository interface {
	GetOrder(id int) (models.Order, error)
	GetOrders(page, pageSize int, sortBy, direction string) ([]models.Order, error)
	GetTotalOrders() (int, error)
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func OrderRepositoryInit(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{db: db}
}

func (o OrderRepositoryImpl) GetOrder(id int) (models.Order, error) {
	var order models.Order
	err := o.db.First(&order, id).Error
	if err != nil {
		log.Error("Error getting order with id: ", id, " in OrderRepositoryImpl. Error: ", err)
		return models.Order{}, err
	}
	return order, nil
}

func (o OrderRepositoryImpl) GetOrders(page, pageSize int, sortBy, direction string) ([]models.Order, error) {
	var orders []models.Order
	orderByStr := sortBy + " " + direction
	err := o.db.Limit(pageSize).Offset((page - 1) * pageSize).Order(orderByStr).Find(&orders).Error
	if err != nil {
		log.Error("Error getting orders with page: ", page, " and pageSize: ", pageSize, " with sortBy: ", sortBy, " and direction: ", direction, " in OrderRepositoryImpl. Error: ", err)
		return nil, err
	}
	return orders, nil
}

func (o OrderRepositoryImpl) GetTotalOrders() (int, error) {
	var totalOrders int64
	err := o.db.Model(&models.Order{}).Count(&totalOrders).Error
	if err != nil {
		log.Error("Error getting total orders in OrderRepositoryImpl. Error: ", err)
		return 0, err
	}
	return int(totalOrders), nil
}
