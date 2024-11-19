package db

import (
	"fmt"
	"os"
	"time"

	"github.com/febriaricandra/book-shop/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DatabaseConnection() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	return nil
}

func DatabaseSeeding() {
	var count int64
	DB.Model(&models.Book{}).Count(&count)
	if count > 0 {
		return
	}

	tx := DB.Begin()
	if tx.Error != nil {
		return
	}

	books := []models.Book{
		{
			Title:       "The Great Gatsby",
			Description: "The Great Gatsby is a 1925 novel by American writer F. Scott Fitzgerald.",
			Category:    "Novel",
			Trending:    true,
			CoverImage:  "https://images-na.ssl-images-amazon.com/images/I/51Zymoq7UnL._AC_SY400_.jpg",
			OldPrice:    10.99,
			NewPrice:    9.99,
		},
		{
			Title:       "To Kill a Mockingbird",
			Description: "To Kill a Mockingbird is a novel by Harper Lee published in 1960.",
			Category:    "Novel",
			Trending:    true,
			CoverImage:  "https://images-na.ssl-images-amazon.com/images/I/51Zymoq7UnL._AC_SY400_.jpg",
			OldPrice:    10.99,
			NewPrice:    9.99,
		},
	}

	for _, book := range books {
		tx.Create(&book)
	}

	tx.Commit()

}
