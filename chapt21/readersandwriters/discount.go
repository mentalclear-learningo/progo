package main

import "encoding/json"

// type DiscountedProduct struct {
// 	*Product
// 	Discount float64
// }

// Updated:
type DiscountedProduct struct {
	*Product `json:",omitempty"`
	Discount float64 `json:",string"`
}

func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error) {
	if dp.Product != nil {
		m := map[string]interface{}{
			"product": dp.Name,
			"cost":    dp.Price - dp.Discount,
		}
		jsn, err = json.Marshal(m)
	}
	return
}

type DiscountedProductTag struct {
	*Product `json:"product"`
	Discount float64
}
