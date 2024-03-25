package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"yab-explorer/models"
	"yab-explorer/services"

	"github.com/gin-gonic/gin"
)

const (
	firstPage = 1
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
	orderIdStr := c.Param("orderId")

	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order Id"})
		return
	}

	order, err := o.service.GetOrder(orderId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (o OrderControllerImpl) GetOrders(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	sort := c.DefaultQuery("sort", "order_id")
	direction := c.DefaultQuery("direction", "desc")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageSize number"})
		return
	}

	if page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page must be greater than 0"})
		return
	}

	if pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pageSize must be greater than 0"})
		return
	}

	if !models.SortArrayContains(sort) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort parameter"})
		return
	}

	if !models.DirectionArrayContains(direction) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid direction parameter"})
		return
	}

	paginatedSearchResult, err := o.service.GetOrders(page, pageSize, sort, direction)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}

	if paginatedSearchResult.PageCount == 0 {
		c.JSON(http.StatusNoContent, gin.H{"error": "No orders found"})
		return
	}

	addLinkHeader(c, paginatedSearchResult)

	c.JSON(http.StatusOK, paginatedSearchResult.Results)
}

func addLinkHeader(c *gin.Context, paginatedSearchResult models.PaginatedSearchResult) {
	page := paginatedSearchResult.Page
	pageSize := paginatedSearchResult.PageSize
	totalOrders := paginatedSearchResult.PageCount

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	baseURL := scheme + "://" + c.Request.Host + c.Request.URL.Path
	baseURL = strings.TrimSuffix(baseURL, "/")
	nextPage := page + 1
	prevPage := page - 1

	lastPage := int(math.Ceil(float64(totalOrders) / float64(pageSize)))

	var linkHeader strings.Builder

	if page < lastPage {
		linkHeader.WriteString(fmt.Sprintf("<%s?page=%d&pageSize=%d>; rel=\"next\"", baseURL, nextPage, pageSize))
	}

	if page > firstPage {
		if linkHeader.Len() > 0 {
			linkHeader.WriteString(", ")
		}
		linkHeader.WriteString(fmt.Sprintf("<%s?page=%d&pageSize=%d>; rel=\"prev\"", baseURL, prevPage, pageSize))
	}

	if linkHeader.Len() > 0 {
		linkHeader.WriteString(", ")
	}

	linkHeader.WriteString(fmt.Sprintf("<%s?page=%d&pageSize=%d>; rel=\"first\"", baseURL, firstPage, pageSize))
	linkHeader.WriteString(", ")
	linkHeader.WriteString(fmt.Sprintf("<%s?page=%d&pageSize=%d>; rel=\"last\"", baseURL, lastPage, pageSize))

	linkHeaderStr := linkHeader.String()

	c.Header("Link", linkHeaderStr)

}
