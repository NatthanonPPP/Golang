package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("Connect successful.")

	db.AutoMigrate(&Book{})
	fmt.Println("Migrate successful")

	// Create
	// newBook := &Book{
	// 	Name:        "Phim",
	// 	Author:      "PPP",
	// 	Description: "TEST",
	// 	Price:       200,
	// }
	// createBook(db, newBook)

	// Get
	// currentBook := getBook(db, 2)
	// fmt.Println(currentBook)

	// Update
	// currentBook := getBook(db, 2)
	// currentBook.Name = "oeoe"
	// currentBook.Price = 212121
	// updateBook(db, currentBook)
	// fmt.Println(currentBook)

	// Delete
	// deleteBook(db, 1)
}
