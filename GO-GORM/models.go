package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string
	Author      string
	Description string
	Price       uint
}

func createBook(db *gorm.DB, book *Book) {
	result := db.Create(book)
	if result.Error != nil {
		log.Fatalf("Error creating book: %v", result.Error)
	}

	fmt.Println("Create Book Successful")
}

func getBook(db *gorm.DB, id uint) *Book {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		log.Fatalf("Error get Book: %v", result.Error)
	}

	return &book
}

func updateBook(db *gorm.DB, book *Book) {
	result := db.Save(&book)
	if result.Error != nil {
		log.Fatalf("Update book failed: %v", result.Error)
	}

	fmt.Println("Update Book successful")
}

func deleteBook(db *gorm.DB, id uint) {
	var book Book

	// Soft Delete
	// result := db.Delete(&book, id)

	// Delete permanently
	result := db.Unscoped().Delete(&book, id)

	if result.Error != nil {
		log.Fatalf("Delete book failed: %v", result.Error)
	}

	fmt.Println("Delete Book successful")
}
