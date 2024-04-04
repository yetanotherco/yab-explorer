package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"yab-explorer/domain/dtos"
	"yab-explorer/domain/models"
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

// GetOrders godoc
//
//	@Summary		Get orders
//	@Description	Get orders
//	@Tags			orders
//	@Produce		json
//	@Param			page		query		int		false	"Page number"
//	@Param			pageSize	query		int		false	"Page size"
//	@Param			sort		query		string	false	"Sort by"
//	@Param			direction	query		string	false	"Sort direction"
//	@Success		200			{object}	[]models.Order
//	@Failure		400			{object}	models.HttpError
//	@Failure		404			{object}	models.HttpError
//	@Failure		500			{object}	models.HttpError
//	@Router			/orders [get]
func (o OrderControllerImpl) GetOrders(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	sort := c.DefaultQuery("sort", "order_id")
	direction := c.DefaultQuery("direction", "desc")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		models.NewError(c, http.StatusBadRequest, fmt.Errorf("page must be an integer"))
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		models.NewError(c, http.StatusBadRequest, fmt.Errorf("pageSize must be an integer"))
		return
	}

	if page < 1 {
		models.NewError(c, http.StatusBadRequest, fmt.Errorf("page must be greater than 0"))
		return
	}

	if pageSize < 1 {
		models.NewError(c, http.StatusBadRequest, fmt.Errorf("pageSize must be greater than 0"))
		return
	}

	if !models.SortArrayContains(sort) {
		models.NewError(c, http.StatusBadRequest, fmt.Errorf("invalid sort parameter"))
		return
	}

	if !models.DirectionArrayContains(direction) {
		models.NewError(c, http.StatusBadRequest, fmt.Errorf("invalid direction parameter"))
		return
	}

	paginatedSearchResult, err := o.service.GetOrders(page, pageSize, sort, direction)

	if err != nil {
		models.NewError(c, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	if paginatedSearchResult.PageCount == 0 {
		models.NewError(c, http.StatusNoContent, fmt.Errorf("no orders found"))
		return
	}

	addLinkHeader(c, paginatedSearchResult)

	c.JSON(http.StatusOK, paginatedSearchResult.Results)
}

func addLinkHeader(c *gin.Context, paginatedSearchResult dtos.PaginatedSearchResultDto) {
	page := paginatedSearchResult.Page
	pageSize := paginatedSearchResult.PageSize
	totalOrders := paginatedSearchResult.ResultsCount

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
