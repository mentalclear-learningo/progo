package store

var standardTax = newTaxRate(0.25, 20)

// Product describes an item for sale
type Product struct {
	Name, Category string  // Start with upper, exported
	price          float64 // Start with lower, not exported
}

// Constructor function:
func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) GetPrice() float64 {
	return standardTax.calcTax(p)
}

func (p *Product) SetPrice(newPrice float64) {
	p.price = newPrice
}
