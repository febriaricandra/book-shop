package handlers

import (
	"log/slog"
	"net/http"
	"strconv"
	"sync"

	"github.com/febriaricandra/book-shop/internal/models"
	"github.com/febriaricandra/book-shop/internal/services"
	"github.com/febriaricandra/book-shop/internal/utils"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: service}
}

func (h *OrderHandler) GetOrdersForUser(c *gin.Context) {
	if userId, exists := c.Get("userId"); !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	} else {
		orders, err := h.orderService.GetOrdersForUser(userId.(uint))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, orders)
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order models.Order
	var orderInput struct {
		models.BaseModel
		Name       string          `json:"name" gorm:"type:varchar(255);not null"`
		Email      string          `json:"email" gorm:"type:varchar(255);not null"`
		Address    models.Address  `json:"address" gorm:"embedded"`
		Phone      string          `json:"phone" gorm:"type:varchar(20);not null"`
		TotalPrice float64         `json:"total_price" gorm:"column:total_price;not null"`
		UserId     uint            `json:"user_id" gorm:"not null"`
		BookIds    []int           `json:"book_ids"`
		Shipping   models.Shipping `json:"shipping" gorm:"embedded"`
	}

	if err := c.ShouldBindJSON(&orderInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userId, exists := c.Get("userId"); !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	} else {
		order.UserId = userId.(uint)
	}

	order.Name = orderInput.Name
	order.Email = orderInput.Email
	order.Address = orderInput.Address
	order.Shipping = orderInput.Shipping
	slog.Info("Order name", "name", order.Shipping)
	slog.Info("Order address", "address", order.Address)
	order.Phone = orderInput.Phone
	order.TotalPrice = orderInput.TotalPrice

	// Create the order first
	orderId, err := h.orderService.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	slog.Info("Order created successfully", "order_id", orderId)

	// Initiate wait group
	var wg sync.WaitGroup
	var orderBookErrors []error

	// Iterate over the book ids and create order book
	for _, bookId := range orderInput.BookIds {
		wg.Add(1)
		go func(bookId int) {
			defer wg.Done()
			// Create order book
			err = h.orderService.CreateOrderBook(orderId, uint(bookId))
			if err != nil {
				orderBookErrors = append(orderBookErrors, err)
				slog.Error("Failed to create order book", "error", err.Error())
			}
		}(bookId)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Check for errors in creating order books
	if len(orderBookErrors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Some order books could not be created", "details": orderBookErrors})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetOrderById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order id"})
		return
	}

	order, err := h.orderService.GetOrderById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	//convert the page and page size to int
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number", "status": false})
		return
	}

	page_size, err := strconv.Atoi(pageSizeStr)
	if err != nil || page_size < 1 || page_size > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size", "status": false})
		return
	}

	orders, totalOrders, err := h.orderService.GetAllOrders(page, page_size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": false})
		return
	}

	totalItems, totalPages := utils.CalculatePagination(totalOrders, page, page_size)

	c.JSON(http.StatusOK, gin.H{"data": orders, "page": page, "page_size": page_size, "total_items": totalItems, "total_pages": totalPages, "status": true})
}
