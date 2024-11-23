package main

import (
	"fmt"
	"log/slog"

	"time"

	"github.com/febriaricandra/book-shop/internal/handlers"
	"github.com/febriaricandra/book-shop/internal/models"
	"github.com/febriaricandra/book-shop/internal/repositories"
	"github.com/febriaricandra/book-shop/internal/routers"
	"github.com/febriaricandra/book-shop/internal/services"
	"github.com/febriaricandra/book-shop/pkg/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Error loading .env file")
	}

	if err := db.DatabaseConnection(); err != nil {
		slog.Error("Error connecting to database", "error", err)
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	// Migrate the schema
	err = db.DB.AutoMigrate(&models.User{}, &models.Book{}, &models.Order{}, &models.OrderBook{})

	if err != nil {
		slog.Error("Error migrating models", "error", err)
		panic(fmt.Sprintf("failed to migrate models: %v", err))
	}

	db.DatabaseSeeding()
}

func main() {

	// Initialize repositories and services
	orderRepo := repositories.NewOrderRepository(db.DB)
	bookRepo := repositories.NewBookRepository(db.DB)
	userRepo := repositories.NewUserRepository(db.DB)

	orderService := services.NewOrderService(orderRepo)
	bookService := services.NewBookService(bookRepo)
	userService := services.NewUserService(userRepo)

	// Initialize handlers
	orderHandler := handlers.NewOrderHandler(orderService)
	bookHandler := handlers.NewBookHandler(bookService)
	userHandler := handlers.NewUserHandler(userService)

	// entry point of the application
	router := gin.Default()
	router.Static("/uploads", "./uploads")

	//CORs configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*", "Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//init route
	routers.BookRouter(router, bookHandler)
	routers.UserRouter(router, userHandler)
	routers.OrderRouter(router, orderHandler)

	router.Use(gin.Logger())

	err := router.Run(":8080")
	if err != nil {
		slog.Error("Error starting server", "error", err)
		panic(err)
	}
}
