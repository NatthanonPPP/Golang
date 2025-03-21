package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

var db *sql.DB

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	db = sdb

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/product/:id", getProductHandler)

	app.Listen(":8080")

	// fmt.Println("Connect Database Successful.")

	// err = createProduct(&Product{Name: "PPP", Price: 200})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Create Successful.")

	// product, err := getProduct(2)
	// fmt.Println("Get Successful.", product)

	// product, err := updateProduct(2, &Product{Name: "ooPP", Price: 200})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Update Products Successful.", product)

	// err = deleteProduct(7)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Delete Product Successful.")

	// products, err := getProducts()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(products)
}

func getProductHandler(c *fiber.Ctx) error {
	return c.SendString("")
}
