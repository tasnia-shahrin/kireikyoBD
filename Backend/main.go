package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//---handling CORS----
func handleCORS(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Methods",",GET,POST,PUT,PATCH,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type")
	w.Header().Set("Content-Type","application/json")
}
//----handling preflight requests-----
func handlePreflightReq(w http.ResponseWriter,r *http.Request){
	if r.Method=="OPTIONS"{
		w.WriteHeader(200)
		return
	}
}

//Sending data
func sendData(w http.ResponseWriter,data interface{},statusCode int){
	w.WriteHeader(statusCode)
	encoder:=json.NewEncoder(w)
	encoder.Encode(data)
}
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
var productList []Product

//--------creating getProducts API-------
func getProducts(w http.ResponseWriter, r *http.Request){

	//---CORS HANDLING---
	handleCORS(w)
	handlePreflightReq(w,r)

	if r.Method!="GET"{
		http.Error(w,"Please give me GET request",400)
		return
	}

	sendData(w,productList,200)
	

}

//---creating creatproducts api----
func createProducts(w http.ResponseWriter,r *http.Request){
	//---CORS HANDLING---
	handleCORS(w)
	//handling preflight requests:
	handlePreflightReq(w,r)
	
	if r.Method!="POST"{
		http.Error(w,"Please give POST method",400)
		return
	}

	/*
	1.take body information from r.Body
	2.create an instance using Product struct with body information
	3.append the instance into productList
	*/
	var newProduct Product //creating instance
	decoder:= json.NewDecoder(r.Body)
	err:=decoder.Decode(&newProduct)
	if err!=nil{
		fmt.Println(err)
		http.Error(w,"please give valid json",400)
		return
	}
	//assign id
	newProduct.ID = len(productList)+1
	productList = append(productList, newProduct)

	sendData(w,newProduct,201)
	 
}
func main(){
	//-----creating router-----
	mux:=http.NewServeMux()

	//---routes for getting products---
	mux.HandleFunc("/products",getProducts)

	//----routes for creating products-----
	mux.HandleFunc("/create-products",createProducts)

	fmt.Println("Server running on port: 8080")
	
	err:=http.ListenAndServe(":8080",mux)
	if err!=nil{
		fmt.Println("Error occurs while starting the server",err)
	}

}

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
	productList = append(productList, prd1)
	productList = append(productList, prd2)
}