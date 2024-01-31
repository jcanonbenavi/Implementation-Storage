package loader

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CustomerJSON struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Condition int    `json:"condition"`
}

type ProductJSON struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type InvoiceJSON struct {
	Datetime   string  `json:"datetime"`
	CustomerId int     `json:"customer_id"`
	Total      float64 `json:"total"`
}

type SalesJSON struct {
	Quantity  int `json:"quantity"`
	InvoiceId int `json:"invoice_id"`
	ProductId int `json:"product_id"`
}

func LoadDataForCustomers(dbConn *sql.DB, path string) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	var customers []CustomerJSON
	if err := json.Unmarshal(fileContent, &customers); err != nil {
		fmt.Println(err)
	}
	if err := dbConn.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	for _, customer := range customers {
		_, err := dbConn.Exec(
			"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
			customer.FirstName, customer.LastName, customer.Condition,
		)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func LoadDataForProducts(dbConn *sql.DB, path string) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	var products []ProductJSON
	if err := json.Unmarshal(fileContent, &products); err != nil {
		fmt.Println(err)
	}
	if err := dbConn.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	for _, product := range products {
		_, err := dbConn.Exec(
			"INSERT INTO products (`description`, `price`) VALUES (?, ?)",
			product.Description, product.Price,
		)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func LoadDataForInvoice(dbConn *sql.DB, path string) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	var invoices []InvoiceJSON
	if err := json.Unmarshal(fileContent, &invoices); err != nil {
		fmt.Println(err)
	}
	if err := dbConn.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	for _, sale := range invoices {
		_, err := dbConn.Exec(
			"INSERT INTO invoices (`datetime`, `customer_id`, `total`) VALUES (?, ?, ?)",
			sale.Datetime, sale.CustomerId, sale.Total,
		)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func LoadDataForSales(dbConn *sql.DB, path string) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	var sales []SalesJSON
	if err := json.Unmarshal(fileContent, &sales); err != nil {
		fmt.Println(err)
	}
	if err := dbConn.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	for _, sale := range sales {
		_, err := dbConn.Exec(
			"INSERT INTO sales (`quantity`, `invoice_id`, `product_id`) VALUES (?, ?, ?)",
			sale.Quantity, sale.InvoiceId, sale.ProductId,
		)
		if err != nil {
			fmt.Println(err)
		}
	}
}
