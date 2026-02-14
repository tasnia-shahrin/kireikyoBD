package product

//---creating product struct---
type Product struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Category string `json:"category"`
	Image string `json:"image"`
	NewPrice float64 `json:"new_price"`
	OldPrice float64 `json:"old_price"`
}
