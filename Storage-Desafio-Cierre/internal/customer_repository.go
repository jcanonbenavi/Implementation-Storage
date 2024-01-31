package internal

// RepositoryCustomer is the interface that wraps the basic methods that a customer repository should implement.
type RepositoryCustomer interface {
	// FindAll returns all customers saved in the database.
	FindAll() (c []Customer, err error)
	// Save saves a customer into the database.
	Save(c *Customer) (err error)
	//GetTop5 returns the top 5 customers
	GetTop5() (c []CustomerTop5, err error)
	//GetByCondition returns the amount of sales by condition of the customer
	GetByCondition() (c []Condition, err error)
}
