package internal

// ServiceCustomer is the interface that wraps the basic methods that a customer service should implement.
type ServiceCustomer interface {
	// FindAll returns all customers
	FindAll() (c []Customer, err error)
	// Save saves a customer
	Save(c *Customer) (err error)
	//GetTop5 returns the top 5 customers
	GetTop5() (c []CustomerTop5, err error)
	//GetCondition returns the condition of the customer
	GetByCondition() (c []Condition, err error)
}
