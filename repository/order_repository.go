package repository

import (
	"yab-explorer/models"

	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

type OrderRepository interface {
	GetOrder(id int) (models.Order, error)
	GetOrders(page, limit int) ([]models.Order, error)
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func OrderRepositoryInit(db *gorm.DB) *OrderRepositoryImpl {
	db.AutoMigrate(&models.Order{})
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

func (o OrderRepositoryImpl) GetOrders(page, limit int) ([]models.Order, error) {
	var orders []models.Order
	err := o.db.Limit(limit).Offset((page - 1) * limit).Order("created_at desc").Find(&orders).Error
	if err != nil {
		log.Error("Error getting orders with page: ", page, " and limit: ", limit, " in OrderRepositoryImpl. Error: ", err)
		return nil, err
	}
	return orders, nil
}
