package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/febriaricandra/book-shop/internal/models"
	"github.com/febriaricandra/book-shop/internal/services"
	"github.com/febriaricandra/book-shop/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookHandler struct {
	bookService *services.BookService
}

func NewBookHandler(bookService *services.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (h *BookHandler) HomeBooks(c *gin.Context) {
	//Get the page and page size from the query string
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	//convert the page and page size to int
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number", "status": false})
		return
	}

	page_size, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size", "status": false})
		return
	}

	topSellerBooks, recommendedBooks, err := h.bookService.GetHomeBooks(page, page_size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"topSellerBooks": topSellerBooks, "recommendedBooks": recommendedBooks, "status": true})
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books, "status": true})
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book
	book.Title = c.PostForm("title")
	book.Description = c.PostForm("description")
	book.Category = c.PostForm("category")
	book.Trending = c.PostForm("trending") == "true"
	book.OldPrice, _ = strconv.ParseFloat(c.PostForm("old_price"), 64)
	book.NewPrice, _ = strconv.ParseFloat(c.PostForm("new_price"), 64)

	//hamdle file upload
	file, err := c.FormFile("cover_image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}

	//Generate a unique file name using uuid and keep the original extension
	extension := filepath.Ext(file.Filename) //get the file extension
	newFileName := fmt.Sprintf("%s%s", uuid.New().String(), extension)

	filePath := filepath.Join("uploads", newFileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}

	book.CoverImage = filePath

	if err := db.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Book created successfully", "data": book})
}

func (h *BookHandler) GetBookById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}

	book, err := h.bookService.GetBookById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "data": book})
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}
	book, err := h.bookService.GetBookById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Book ID", "status": false})
		return
	}

	book.ID = uint(id)
	book.Title = c.PostForm("title")
	book.Description = c.PostForm("description")
	book.Category = c.PostForm("category")
	book.Trending = c.PostForm("trending") == "true"
	book.OldPrice, _ = strconv.ParseFloat(c.PostForm("old_price"), 64)
	book.NewPrice, _ = strconv.ParseFloat(c.PostForm("new_price"), 64)

	//hamdle file upload
	file, err := c.FormFile("cover_image")
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			//No file was uploaded
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
			return
		}
	} else {
		//Generate a unique file name using uuid and keep the original extension
		extension := filepath.Ext(file.Filename) //get the file extension
		newFileName := fmt.Sprintf("%s%s", uuid.New().String(), extension)

		filePath := filepath.Join("uploads", newFileName)

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
			return
		}

		book.CoverImage = filePath
	}

	if err := db.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Book updated successfully", "data": book})
}
