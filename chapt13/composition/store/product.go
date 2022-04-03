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

// Added with Listing 13-23
type Describable interface {
	GetName() string
	GetCategory() string
	// Composing interfaces
	ItemForSale // interfaces can enclose another
}

func (p *Product) GetName() string {
	return p.Name
}
func (p *Product) GetCategory() string {
	return p.Category
}
