package handlers

import (
	"log/slog"
	"net/http"
	"strconv"
	"sync"

	"github.com/febriaricandra/book-shop/internal/models"
	"github.com/febriaricandra/book-shop/internal/services"
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
		Name       string         `json:"name" gorm:"type:varchar(255);not null"`
		Email      string         `json:"email" gorm:"type:varchar(255);not null"`
		Address    models.Address `json:"address" gorm:"embedded"`
		Phone      string         `json:"phone" gorm:"type:varchar(20);not null"`
		TotalPrice float64        `json:"total_price" gorm:"column:total_price;not null"`
		UserId     uint           `json:"user_id" gorm:"not null"`

		BookIds []int `json:"book_ids"`
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
	order.Phone = orderInput.Phone
	order.TotalPrice = orderInput.TotalPrice
	order.UserId = orderInput.UserId

	orderId, err := h.orderService.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	slog.Info("Order created successfully", "order_id", orderId)
	//initiate wait group
	var wg sync.WaitGroup

	// iterate over the book ids and create order book
	for _, bookId := range orderInput.BookIds {
		//increment the wait group counter
		wg.Add(1)
		//create goroutine for each book
		go func(bookId int) {
			//decrement the wait group counter when the goroutine completes
			defer wg.Done()
			//create order book
			err = h.orderService.CreateOrderBook(orderId, uint(bookId))
			if err != nil {
				slog.Error("Failed to create order book", "error", err.Error())
			}
		}(bookId)
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
