package handlers

import (
	"Backend/Database"
	"Backend/util"
	"net/http"
)

//--------creating getProducts API-------
func GetProducts(w http.ResponseWriter, r *http.Request){

	util.SendData(w,Database.ProductList,200)

}
