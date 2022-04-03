package store

type ItemForSale interface {
	GetPrice(taxRate float64) float64
}
