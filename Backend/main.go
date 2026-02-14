package main

import (
	"fmt"
	"net/http"
	"Backend/product"
	"Backend/handlers"
	"Backend/globalRouter"
)

//---handling CORS----
// func handleCORS(w http.ResponseWriter){
// 	w.Header().Set("Access-Control-Allow-Origin","*")
// 	w.Header().Set("Access-Control-Allow-Methods",",GET,POST,PUT,PATCH,DELETE,OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers","Content-Type")
// 	w.Header().Set("Content-Type","application/json")
// }
// //----handling preflight requests-----
// func handlePreflightReq(w http.ResponseWriter,r *http.Request){
// 	if r.Method=="OPTIONS"{
// 		w.WriteHeader(200)
// 		return
// 	}
// }


func main(){
	//-----creating router-----
	mux:=http.NewServeMux()

	
	//---routes for getting products---
	mux.Handle("GET /products",http.HandlerFunc(handlers.GetProducts))

	//----routes for creating products-----
	// mux.HandleFunc("/create-products",createProducts)-->previous routing
	mux.Handle("POST /create-products",http.HandlerFunc(handlers.CreateProducts)) //advanced routing

	fmt.Println("Server running on port: 8080")
	
	globalrouter:= globalRouter.GlobalRouter(mux)
	err:=http.ListenAndServe(":8080",globalrouter)
	if err!=nil{
		fmt.Println("Error occurs while starting the server",err)
	}

}

func init(){
	prd1:= product.Product{
		ID:1,
		Name: "Anua Niacinamide 10% + TXA 4% Serum – 30ml",
    	Category: "k-beauty",
    	Image: "https://koreanmartbd.com/wp-content/uploads/2024/05/Anua-Niacinamide-10-TXA-4-Dark-Spot-Correcting-Serum-30ml-1.jpg",
    	NewPrice: 1879.00,
    	OldPrice: 2730,
	}
	prd2:= product.Product{
		ID:2,
		Name: "NIVEA UV Super Water Gel EX SPF50+ PA++++ 80g",
    	Category: "j-beauty",
    	Image: "https://cdn.kireibd.com/storage/all/NIVEA-UV-Super-Water-Gel-EX-SPF50-PA--80g.png",
    	NewPrice: 1380,
    	OldPrice: 1410,
	}
	product.ProductList = append(product.ProductList, prd1)
	product.ProductList = append(product.ProductList, prd2)
}