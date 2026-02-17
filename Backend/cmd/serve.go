package cmd

import (
	"fmt"
	"net/http"
	"Backend/handlers"
	"Backend/globalRouter"
)

func Serve(){
	//-----creating router-----
	mux:=http.NewServeMux()

	
	//---routes for getting products---
	mux.Handle("GET /products",http.HandlerFunc(handlers.GetProducts))

	//----routes for creating products-----
	// mux.HandleFunc("/create-products",createProducts)-->previous routing
	mux.Handle("POST /products",http.HandlerFunc(handlers.CreateProducts)) //advanced routing

	mux.Handle("GET /products/{id}",http.HandlerFunc(handlers.GetProductByID))

	fmt.Println("Server running on port: 8080")
	
	globalrouter:= globalRouter.GlobalRouter(mux)
	err:=http.ListenAndServe(":8080",globalrouter)
	if err!=nil{
		fmt.Println("Error occurs while starting the server",err)
	}

}