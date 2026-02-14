package main

import(
	"fmt"
	"net/http"
	"encoding/json"
)



//---creating createproducts api----
func createProducts(w http.ResponseWriter,r *http.Request){

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
