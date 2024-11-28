package routers

import (
	"github.com/febriaricandra/book-shop/internal/handlers"
	"github.com/febriaricandra/book-shop/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine, h *handlers.BookHandler) {

	//public route v1
	public := router.Group("/api")
	{
		public.GET("/books", h.GetBooks)
		public.GET("/books/:id", h.GetBookById)
		public.GET("/books/home", h.HomeBooks)
	}

	//private route v1
	private := router.Group("/api")
	private.Use(middlewares.AuthMiddleware())
	{
		private.POST("/books", middlewares.AdminMiddleware(), h.CreateBook)
		private.PUT("/books/:id", middlewares.AdminMiddleware(), h.UpdateBook)
	}
}

func UserRouter(router *gin.Engine, h *handlers.UserHandler) {
	public := router.Group("/api")
	{
		public.POST("/register", h.RegisterUser)
		public.POST("/login", h.Login)
		public.POST("/refresh", h.Refresh)
	}

	private := router.Group("/api")
	private.Use(middlewares.AuthMiddleware())
	{
		private.GET("/profile", h.Profile)
	}
}

func OrderRouter(router *gin.Engine, h *handlers.OrderHandler) {
	private := router.Group("/api")
	private.Use(middlewares.AuthMiddleware())
	{
		private.POST("/orders", h.CreateOrder)
		private.GET("/orders/:id", h.GetOrderById)
		private.GET("/user-orders", h.GetOrdersForUser)
	}
}
