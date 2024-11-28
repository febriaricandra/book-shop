package middlewares

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/febriaricandra/book-shop/internal/repositories"
	"github.com/febriaricandra/book-shop/internal/services"
	"github.com/febriaricandra/book-shop/pkg/db"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		bearerToken := ""

		if len(strings.Split(token, " ")) == 2 {
			bearerToken = strings.Split(token, " ")[1]
		}

		if bearerToken == "" {
			slog.Error("No bearer token found")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No bearer token found", "status": false})
			c.Abort()
			return
		}

		userRepo := repositories.NewUserRepository(db.DB)
		userService := services.NewUserService(userRepo)

		claims, err := userService.VerifyToken(bearerToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		user, err := userRepo.GetUserByEmail(claims.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		if user.Email != claims.Email {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		info := map[string]interface{}{
			"email":   user.Email,
			"name":    user.Name,
			"isAdmin": user.IsAdmin,
		}

		c.Set("Claims", claims)
		c.Set("userId", user.ID)
		c.Set("email", user.Email)
		c.Set("isAdmin", user.IsAdmin)
		c.Set("info", info)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, exists := c.Get("isAdmin")
		if !exists || !isAdmin.(bool) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access Forbidden", "status": false})
			c.Abort()
			return
		}
		c.Next()
	}
}
