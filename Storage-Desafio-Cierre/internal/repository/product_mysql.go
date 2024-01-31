package repository

import (
	"database/sql"

	"app/internal"
)

// NewProductsMySQL creates new mysql repository for product entity.
func NewProductsMySQL(db *sql.DB) *ProductsMySQL {
	return &ProductsMySQL{db}
}

// ProductsMySQL is the MySQL repository implementation for product entity.
type ProductsMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all products from the database.
func (r *ProductsMySQL) FindAll() (p []internal.Product, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `description`, `price` FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var pr internal.Product
		// scan the row into the product
		err := rows.Scan(&pr.Id, &pr.Description, &pr.Price)
		if err != nil {
			return nil, err
		}
		// append the product to the slice
		p = append(p, pr)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the product into the database.
func (r *ProductsMySQL) Save(p *internal.Product) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO products (`description`, `price`) VALUES (?, ?)",
		(*p).Description, (*p).Price,
	)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*p).Id = int(id)

	return
}

// GetTop5 returns the top 5 products and the total of sales for each one.
func (r *ProductsMySQL) GetTop5() (p []internal.ProductsTop5, err error) {
	rows, err := r.db.Query("SELECT p.description, sum(s.quantity) as total FROM sales s INNER JOIN products p ON s.product_id = p.id GROUP BY s.product_id ORDER BY total DESC LIMIT 5;")
	if err != nil {
		return nil, err
	}
	//iterate over the rows and append to the slice
	for rows.Next() {
		//initialize the product top5
		var top5 internal.ProductsTop5
		// scan the row into the product top5
		err := rows.Scan(&top5.Description, &top5.Total)
		if err != nil {
			return nil, err
		}
		// append the product to the slice
		p = append(p, top5)

	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}
