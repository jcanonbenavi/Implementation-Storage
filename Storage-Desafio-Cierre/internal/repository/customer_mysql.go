package repository

import (
	"database/sql"
	"fmt"

	"app/internal"
)

// NewCustomersMySQL creates new mysql repository for customer entity.
func NewCustomersMySQL(db *sql.DB) *CustomersMySQL {
	return &CustomersMySQL{db}
}

// CustomersMySQL is the MySQL repository implementation for customer entity.
type CustomersMySQL struct {
	// db is the database connection.
	db *sql.DB
}

func (r *CustomersMySQL) LoadData(customer internal.CustomerAttributes) (c *internal.Customer, err error) {
	res, err := r.db.Exec(
		"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
		customer.FirstName, customer.LastName, customer.Condition,
	)
	if err != nil {
		fmt.Println(err)
	}
	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}

	c = &internal.Customer{
		CustomerAttributes: customer,
		Id:                 int(id),
	}

	return
}

// FindAll returns all customers from the database.
func (r *CustomersMySQL) FindAll() (c []internal.Customer, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `first_name`, `last_name`, `condition` FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs internal.Customer
		// scan the row into the customer
		err := rows.Scan(&cs.Id, &cs.FirstName, &cs.LastName, &cs.Condition)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the customer into the database.
func (r *CustomersMySQL) Save(c *internal.Customer) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
		(*c).FirstName, (*c).LastName, (*c).Condition,
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
	(*c).Id = int(id)

	return
}

// GetTop5 returns the top 5 customers
func (r *CustomersMySQL) GetTop5() (c []internal.CustomerTop5, err error) {
	rows, err := r.db.Query("SELECT c.first_name, c.last_name, sum(total) amount FROM invoices i INNER JOIN customers c ON i.customer_id = c.id group by i.customer_id order by amount desc LIMIT 5;")
	if err != nil {
		return nil, err
	}
	//iterate over the rows and append to the slice
	for rows.Next() {
		//initialize the customer top5
		var top5 internal.CustomerTop5
		// scan the row into the customer top5
		err := rows.Scan(&top5.FirstName, &top5.LastName, &top5.Amount)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, top5)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}

func (r *CustomersMySQL) GetByCondition() (c []internal.Condition, err error) {
	query := "SELECT c.condition, sum(total) amount " +
		"FROM invoices i INNER JOIN customers c ON i.customer_id = c.id " +
		"group by c.condition order by amount;"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	//iterate over the rows and append to the slice
	for rows.Next() {
		//initialize the condition
		var condition internal.Condition
		// scan the row into the condition
		err := rows.Scan(&condition.Condition, &condition.Total)
		if err != nil {
			return nil, err
		}
		// append the condition to the slice
		c = append(c, condition)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
