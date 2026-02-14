package main

import "net/http"


//--------creating getProducts API-------
func getProducts(w http.ResponseWriter, r *http.Request){

	sendData(w,productList,200)

}
