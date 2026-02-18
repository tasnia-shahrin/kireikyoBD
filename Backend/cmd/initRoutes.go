package cmd

import (
	"Backend/middleware"
	"net/http"
	"Backend/handlers"
)

func initRoutes(mux *http.ServeMux,manager *middleware.Manager){
	//---routes for getting products---
	mux.Handle(
		"GET /products",
		
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)

	//----routes for creating products-----
	// mux.HandleFunc("/create-products",createProducts)-->previous routing
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProducts),		
		),
	) //advanced routing

	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProductByID),
		),
	)

}