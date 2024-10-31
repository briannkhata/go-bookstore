package models

import (
	"log"

	"github.com/briannkhata/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	if db == nil {
		log.Fatal("Database connection failed")
	}
	if err := db.AutoMigrate(&Book{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
}

// CreateBook creates a new book record
func (b *Book) CreateBook() *Book {
	if db == nil {
		log.Println("Database connection is null")
		return nil
	}
	if err := db.Create(b).Error; err != nil {
		log.Printf("Error creating book record: %v", err)
		return nil
	}
	return b
}

// GetAllBooks retrieves all book records
func GetAllBooks() []Book {
	var books []Book
	if db == nil {
		log.Println("Database connection is null")
		return books
	}
	db.Find(&books)
	return books
}

// GetBookById retrieves a book by its ID
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	if db == nil {
		log.Println("Database connection is null")
		return nil, nil
	}
	res := db.Where("ID = ?", Id).First(&book)
	return &book, res
}

// DeleteBook deletes a book by its ID
func DeleteBook(Id int64) *Book {
	var book Book
	if db == nil {
		log.Println("Database connection is null")
		return nil
	}
	db.Where("ID = ?", Id).Delete(&book)
	return &book
}
