package repositories

import (
	"log/slog"

	"github.com/febriaricandra/book-shop/internal/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
	GetBookById(bookId uint) (*models.Book, error)
	GetAllBooks(page, pageSize int) ([]models.Book, int, error)
	UpdateBook(book *models.Book) error
	DeleteBook(bookId uint) error
	GetHomeBooks(page, pageSize int) ([]models.Book, []models.Book, int, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) CreateBook(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) GetBookById(bookId uint) (*models.Book, error) {
	var book models.Book

	err := r.db.First(&book, bookId).Error
	return &book, err
}

func (r *bookRepository) GetAllBooks(page, pageSize int) ([]models.Book, int, error) {
	var books []models.Book
	var totalBook int64

	err := r.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&books).Error
	if err != nil {
		slog.Error("Error getting all books", "error", err.Error())
		return nil, 0, err
	}

	err = r.db.Model(&models.Book{}).Count(&totalBook).Error
	if err != nil {

		slog.Error("Error getting total books", "error", err.Error())
		return nil, 0, err
	}

	return books, int(totalBook), nil
}

func (r *bookRepository) UpdateBook(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) DeleteBook(bookId uint) error {
	return r.db.Delete(&models.Book{}, bookId).Error
}

func (r *bookRepository) GetHomeBooks(page, pageSize int) ([]models.Book, []models.Book, int, error) {
	var topSellerBooks []models.Book
	var recommendedBooks []models.Book
	var total int64

	offset := (page - 1) * pageSize
	err := r.db.Where("trending = ?", true).Offset(offset).Limit(pageSize).Find(&topSellerBooks).Error
	if err != nil {
		slog.Error("Error getting top seller books", "error", err.Error())
		return nil, nil, 0, err
	}

	err = r.db.Model(&models.Book{}).Count(&total).Error
	if err != nil {
		slog.Error("Error getting total books", "error", err.Error())
		return nil, nil, 0, err
	}
	if len(topSellerBooks) > 0 {
		// Get recommended books except top seller books
		ids := make([]uint, len(topSellerBooks))
		for i, book := range topSellerBooks {
			ids[i] = book.ID
		}
		err = r.db.Where("trending = ? AND id NOT IN (?)", false, ids).Find(&recommendedBooks).Error
		if err != nil {
			slog.Error("Error getting recommended books", "error", err.Error())
			return topSellerBooks, nil, int(total), err
		}
		return topSellerBooks, recommendedBooks, int(total), nil
	}
	err = r.db.Where("trending = ?", false).Find(&recommendedBooks).Error
	if err != nil {
		slog.Error("Error getting recommended books", "error", err.Error())
		return nil, recommendedBooks, int(total), err
	}
	return topSellerBooks, recommendedBooks, int(total), nil
}
