package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Product struct {
	ProductName  string
	Unit         string
	Price        float64
	CategoryName string
}

func main() {
	connStr := "user=postgres password=vakhaboff dbname=shaxboz sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `
        SELECT p.product_name, p.unit, p.price, c.category_name FROM products p
        JOIN categories c ON p.category_id = c.category_id
        WHERE c.category_name = 'Beverages';
    `
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ProductName, &product.Unit, &product.Price, &product.CategoryName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Mahsulot: %s, Narxi: %f, Miqdori: %s, Kategoriya: %s\n", product.ProductName, product.Price, product.Unit, product.CategoryName)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
