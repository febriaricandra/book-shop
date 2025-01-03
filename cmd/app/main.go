package main

import (
	"fmt"
	"log/slog"
	"os"

	"time"

	cfg "github.com/febriaricandra/book-shop/config"

	"github.com/aws/aws-sdk-go-v2/service/s3"
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

var R2Client *s3.Client

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Error loading .env file")
	}

	R2Client, err = cfg.InitR2Client()
	if err != nil {
		slog.Error("Error initializing R2 client", "error", err)
		panic(fmt.Sprintf("failed to initialize R2 client: %v", err))
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
	bookHandler := handlers.NewBookHandler(bookService, R2Client, "bookshop", os.Getenv("ENDPOINT_URL"))
	userHandler := handlers.NewUserHandler(userService)
	rajaOngkirHandler := handlers.NewRajaOngkirHandler(os.Getenv("RAJAONGKIR_API_KEY"))

	// entry point of the application
	router := gin.Default()
	// router.Static("/uploads", "./uploads")

	//CORs configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"*", "Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	//init route
	routers.BookRouter(router, bookHandler)
	routers.UserRouter(router, userHandler)
	routers.OrderRouter(router, orderHandler)
	routers.RajaOngkirRouter(router, rajaOngkirHandler)

	router.Use(gin.Logger())

	err := router.Run(":8080")
	if err != nil {
		slog.Error("Error starting server", "error", err)
		panic(err)
	}
}
