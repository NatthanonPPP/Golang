package main

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for i, book := range books {
		if bookId == book.ID {
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author
			return c.JSON(books[i])
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func deleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for i, book := range books {
		if bookId == book.ID {
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err = c.SaveFile(file, "./uploads/"+file.Filename)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.SendString("File upload completed.")
}

func testHTML(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, Go!",
	})
}

func getEnv(c *fiber.Ctx) error {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "defaultsecret"
	}
	return c.JSON(fiber.Map{
		"SECRET": secret,
	})
}
