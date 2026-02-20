package handlers

import (
	"Backend/Database"
	
	"Backend/util"
	
	"encoding/json"
	"fmt"
	"net/http"
	
)

//---creating createproducts api----
func CreateProducts(w http.ResponseWriter,r *http.Request){

	
	/*
	1.take body information from r.Body
	2.create an instance using Product struct with body information
	3.append the instance into productList
	*/
	var newProduct Database.Product //creating instance
	decoder:= json.NewDecoder(r.Body)
	err:=decoder.Decode(&newProduct)
	if err!=nil{
		fmt.Println(err)
		http.Error(w,"please give valid json",400)
		return
	}
	
	createdProduct:= Database.Store(newProduct)
	
	util.SendData(w,createdProduct,201)
	
}

