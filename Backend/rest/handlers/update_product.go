package handlers

import (
	"Backend/Database"
	"Backend/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter,r *http.Request){

	productID:=r.PathValue("id")

	pId,err:=strconv.Atoi(productID)
	if err!=nil{
		http.Error(w,"please give me a valid product id",400)
		return
	}
	var newProduct Database.Product //creating instance
	decoder:= json.NewDecoder(r.Body)
	err=decoder.Decode(&newProduct)
	if err!=nil{
		fmt.Println(err)
		http.Error(w,"please give valid json",400)
		return
	}
	newProduct.ID = pId
	Database.Update(newProduct)

	util.SendData(w,"Successfully updated product",201)
}