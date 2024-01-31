package internal

// CustomerAttributes is the struct that represents the attributes of a customer.
type CustomerAttributes struct {
	// FirstName is the first name of the customer.
	FirstName string
	// LastName is the last name of the customer.
	LastName string
	// Condition is the condition of the customer.
	Condition int
}

// Customer is the struct that represents a customer.
type Customer struct {
	// Id is the unique identifier of the customer.
	Id int
	// CustomerAttributes is the attributes of the customer.
	CustomerAttributes
}

// Top5 is the struct that represents the top 5 products.
type CustomerTop5 struct {
	// FirstName is the first name of the customer.
	FirstName string
	// LastName is the last name of the customer.
	LastName string
	//Amount is the total of sales of the customer.
	Amount float64
}

type Condition struct {
	// Condition is the condition of the customer.
	Condition int
	// Total is the total of sales of the customer.
	Total float64
}
