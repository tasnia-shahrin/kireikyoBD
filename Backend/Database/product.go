package Database

//---creating product struct---
type Product struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Category string `json:"category"`
	Image string `json:"image"`
	NewPrice float64 `json:"new_price"`
	OldPrice float64 `json:"old_price"`
}

//--empty slice---
var ProductList []Product

func init(){
	prd1:= Product{
		ID:1,
		Name: "Anua Niacinamide 10% + TXA 4% Serum – 30ml",
    	Category: "k-beauty",
    	Image: "https://koreanmartbd.com/wp-content/uploads/2024/05/Anua-Niacinamide-10-TXA-4-Dark-Spot-Correcting-Serum-30ml-1.jpg",
    	NewPrice: 1879.00,
    	OldPrice: 2730,
	}
	prd2:= Product{
		ID:2,
		Name: "NIVEA UV Super Water Gel EX SPF50+ PA++++ 80g",
    	Category: "j-beauty",
    	Image: "https://cdn.kireibd.com/storage/all/NIVEA-UV-Super-Water-Gel-EX-SPF50-PA--80g.png",
    	NewPrice: 1380,
    	OldPrice: 1410,
	}
	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
}