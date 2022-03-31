package main

type Taxes struct {
	tax       string
	taxAmount float64
}

func (t Taxes) getName() string {
	return t.tax
}

func (t Taxes) getCost(current bool) float64 {
	if current {
		return t.taxAmount
	}
	return 0
}
