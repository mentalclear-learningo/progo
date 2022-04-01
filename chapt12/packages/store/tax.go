package store

const defaultTaxRate float64 = 0.2
const minThreshold = 10

// Problematic starting point:
// var categoryMaxPrices = map[string]float64{
// 	"Watersports": 250 + (250 * defaultTaxRate),
// 	"Soccer":      150 + (150 * defaultTaxRate),
// 	"Chess":       50 + (50 * defaultTaxRate),
// }

var categoryMaxPrices = map[string]float64{
	"Watersports": 250,
	"Soccer":      150,
	"Chess":       50,
}

// Package initialization function, runs automatically on execution
func init() {
	for category, price := range categoryMaxPrices {
		categoryMaxPrices[category] = price + (price * defaultTaxRate)
	}
}

type taxRate struct {
	rate, threshold float64
}

func newTaxRate(rate, threshold float64) *taxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}
	if threshold < minThreshold {
		threshold = minThreshold
	}
	return &taxRate{rate, threshold}
}
func (taxRate *taxRate) calcTax(product *Product) (price float64) {
	if product.price > taxRate.threshold {
		price = product.price + (product.price * taxRate.rate)
	} else {
		price = product.price
	}
	if max, ok := categoryMaxPrices[product.Category]; ok && price > max {
		price = max
	}
	return
}
