package main

import (
	"app/internal/loader"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// env
	// ...

	// app
	// - config
	cfg := mysql.Config{
		User:   "root",
		Passwd: "MayaSparky94",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "fantasy_products",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	// - set up
	CustomersPath := "../docs/db/json/customers.json"
	loader.LoadDataForCustomers(db, CustomersPath)
	ProductsPath := "../docs/db/json/products.json"
	loader.LoadDataForProducts(db, ProductsPath)
	InvoicePath := "../docs/db/json/invoices.json"
	loader.LoadDataForInvoice(db, InvoicePath)
	SalesPath := "../docs/db/json/sales.json"
	loader.LoadDataForSales(db, SalesPath)
}
