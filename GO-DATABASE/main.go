package main

import (
	"database/sql"
	"fmt"
	"log"

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
	ID    int
	Name  string
	Price int
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

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect Database Successful.")

	// err = creatProduct(&Product{Name: "Go product 2", Price: 444})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Create Successful.")

	// product, err := getProduct(2)
	// fmt.Println("Get Successful.", product)

	// product, err := updateProduct(2, &Product{Name: "UUU", Price: 333})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Update Products Successful.", product)

	// err = deleteProduct(7)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Delete Product Successful.")

	products, err := getProducts()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(products)
}

func creatProduct(product *Product) error {
	_, err := db.Exec(
		"INSERT INTO public.products(name, price) VALUES ($1, $2);",
		product.Name,
		product.Price,
	)
	return err
}

func getProduct(id int) (Product, error) {
	var p Product

	row := db.QueryRow(
		"SELECT id, name, price FROM products WHERE id=$1;",
		id,
	)

	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func getProducts() ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}

	var products []Product

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func updateProduct(id int, product *Product) (Product, error) {
	var p Product
	row := db.QueryRow(
		"UPDATE public.products SET name=$1, price=$2 WHERE id=$3 RETURNING id, name, price;",
		product.Name,
		product.Price,
		id,
	)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}
	return p, err
}

func deleteProduct(id int) error {
	_, err := db.Exec(
		"DELETE FROM products WHERE id=$1;",
		id,
	)
	return err
}
