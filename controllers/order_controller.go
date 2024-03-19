package controllers

import (
	"net/http"
	"strconv"
	"yab-explorer/services"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	GetOrder(c *gin.Context)
	GetOrders(c *gin.Context)
}

type OrderControllerImpl struct {
	service services.OrderService
}

func OrderControllerInit(orderService services.OrderService) *OrderControllerImpl {
	return &OrderControllerImpl{service: orderService}
}

func (o OrderControllerImpl) GetOrder(c *gin.Context) {
	orderIDStr := c.Param("orderID")

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := o.service.GetOrder(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (o OrderControllerImpl) GetOrders(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
		return
	}

	if page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page must be greater than 0"})
		return
	}

	if limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Limit must be greater than 0"})
		return
	}

	orders, err := o.service.GetOrders(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}

	c.JSON(http.StatusOK, orders)

}
