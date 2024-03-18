package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"yab-explorer/services"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	GetOrder(ctx *gin.Context)
	GetOrders(ctx *gin.Context)
}

type controller struct {
	service services.OrderService
}

func New(service services.OrderService) OrderController {
	return &controller{service: service}
}

func (c *controller) GetOrder(ctx *gin.Context) {
	orderIDStr := ctx.Param("orderID")

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := c.service.GetOrder(orderID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Order not found"})
		return
	}
	ctx.JSON(http.StatusOK, order)
}

func (c *controller) GetOrders(ctx *gin.Context) {
	fmt.Println("GetOrders")
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
		return
	}

	if page < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Page must be greater than 0"})
		return
	}

	if limit < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Limit must be greater than 0"})
		return
	}

	orders, err := c.service.GetOrders(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}

	ctx.JSON(http.StatusOK, orders)

}
