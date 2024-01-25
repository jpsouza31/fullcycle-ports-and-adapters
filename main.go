package main

import (
	"database/sql"
	"ports-and-adapters/adapters/db"
	"ports-and-adapters/application"
)

func main() {
	Db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db.NewProductDb(Db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Exemplo", 30)

	productService.Enable(product)
}
