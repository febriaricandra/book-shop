package repositories

import (
	"log/slog"

	"github.com/febriaricandra/book-shop/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) (uint, error)
	GetOrderById(id uint) (*models.Order, error)
	GetAllOrders(page, pageSize int) ([]models.Order, int, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(id uint) error
	CreateOrderBook(uint, uint) error
	GetOrdersForUser(uint) ([]models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) CreateOrder(order *models.Order) (uint, error) {
	if err := r.db.Create(order).Error; err != nil {
		return 0, err
	}
	return order.ID, nil
}

func (r *orderRepository) GetOrderById(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Books").First(&order, id).Error
	return &order, err
}

func (r *orderRepository) GetAllOrders(page, pageSize int) ([]models.Order, int, error) {
	var orders []models.Order
	var totalOrders int64

	err := r.db.Preload("Books").Preload("User").Offset((page - 1) * pageSize).Limit(pageSize).Find(&orders).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Model(&models.Order{}).Count(&totalOrders).Error
	if err != nil {
		return nil, 0, err
	}

	return orders, int(totalOrders), nil
}

func (r *orderRepository) UpdateOrder(order *models.Order) error {
	return r.db.Save(order).Error
}

func (r *orderRepository) DeleteOrder(id uint) error {
	return r.db.Delete(&models.Order{}, id).Error
}

func (r *orderRepository) CreateOrderBook(orderId uint, bookId uint) error {
	orderBook := models.OrderBook{OrderID: orderId, BookID: bookId}
	if err := r.db.Create(&orderBook).Error; err != nil {
		// Log the error for debugging
		slog.Error("Failed to create order book", "error", err.Error())
		return err
	}
	return nil
}

func (r *orderRepository) GetOrdersForUser(userId uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id = ?", userId).Preload("Books").Preload("User").Find(&orders).Error
	return orders, err
}
