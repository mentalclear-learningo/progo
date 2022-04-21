package main

type DiscountedProduct struct {
	*Product
	Discount float64
}

type DiscountedProductTag struct {
	*Product `json:"product"`
	Discount float64
}
