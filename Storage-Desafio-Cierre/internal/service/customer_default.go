package service

import "app/internal"

// NewCustomersDefault creates new default service for customer entity.
func NewCustomersDefault(rp internal.RepositoryCustomer) *CustomersDefault {
	return &CustomersDefault{rp}
}

// CustomersDefault is the default service implementation for customer entity.
type CustomersDefault struct {
	// rp is the repository for customer entity.
	rp internal.RepositoryCustomer
}

// FindAll returns all customers.
func (s *CustomersDefault) FindAll() (c []internal.Customer, err error) {
	c, err = s.rp.FindAll()
	return
}

// Save saves the customer.
func (s *CustomersDefault) Save(c *internal.Customer) (err error) {
	err = s.rp.Save(c)
	return
}

// GetTop5 returns the top 5 customers.
func (s *CustomersDefault) GetTop5() (c []internal.CustomerTop5, err error) {
	c, err = s.rp.GetTop5()
	return
}

// GetByCondition returns the amount of sales by condition of the customer
func (s *CustomersDefault) GetByCondition() (c []internal.Condition, err error) {
	c, err = s.rp.GetByCondition()
	return
}
