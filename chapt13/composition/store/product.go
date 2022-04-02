package store

type Product struct {
	Name, Category string
	price          float64
}

// constructor
func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

// Method
func (p *Product) GetPrice(taxRate float64) float64 {
	return p.price + (p.price * taxRate)
}