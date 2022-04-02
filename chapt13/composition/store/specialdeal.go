package store

type SpecialDeal struct {
	Name string
	*Product
	price float64
}

func NewSpecialDeal(name string, p *Product, discount float64) *SpecialDeal {
	return &SpecialDeal{name, p, p.price - discount}
}
func (deal *SpecialDeal) GetDetails() (string, float64, float64) {
	// This is using GetPrice() not on the deal but on the Product
	// because deal.GetPrice() is basically deal.Product.GetPrice()
	return deal.Name, deal.price, deal.GetPrice(0)
}

// If I want to be able to call the Price method and get a result
// that does rely on the SpecialDeal.price field, then I have to define a new method:
func (deal *SpecialDeal) GetPrice(taxRate float64) float64 {
	return deal.price
}

// This stops the promotion of the Product and uses the one from here
