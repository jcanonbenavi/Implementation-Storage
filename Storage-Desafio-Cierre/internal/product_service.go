package internal

// ServiceProduct is the interface that wraps the basic Product methods.
type ServiceProduct interface {
	// FindAll returns all products.
	FindAll() (p []Product, err error)
	// Save saves a product.
	Save(p *Product) (err error)
	//GetTheTop5 returns the top 5 products
	GetTop5() (p []ProductsTop5, err error)
}
