package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func checkMiddleware(c *fiber.Ctx) error {
	fmt.Printf("URL = %s, Method = %s, Time = %s\n", c.OriginalURL(), c.Method(), time.Now())
	return c.Next()
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var memberUser = User{
	Email:    "user@example.com",
	Password: "password123",
}

func login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if user.Email != memberUser.Email || user.Password != memberUser.Password {
		return fiber.ErrUnauthorized
	}
	return c.JSON(fiber.Map{
		"message": "Login success.",
	})
}

func main() {
	// env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Load .env error!!")
	}

	// view template
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	books = append(books, Book{ID: 1, Title: "Water", Author: "Phim"})
	books = append(books, Book{ID: 2, Title: "Drink", Author: "POP"})

	app.Post("/login", login)

	app.Use(checkMiddleware)
	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)
	app.Post("/upload", uploadFile) // upload file

	app.Get("/test-html", testHTML) // view templaate

	app.Get("/config", getEnv) // env

	app.Listen(":8080")
}
