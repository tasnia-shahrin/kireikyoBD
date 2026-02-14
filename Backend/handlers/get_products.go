package handlers

import (
	"Backend/product"
	"Backend/util"
	"net/http"
)

//--------creating getProducts API-------
func GetProducts(w http.ResponseWriter, r *http.Request){

	util.SendData(w,product.ProductList,200)

}
