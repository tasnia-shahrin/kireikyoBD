package rest

import (
	"Backend/rest/middleware"
	"net/http"
	"Backend/rest/handlers"
)

func InitRoutes(mux *http.ServeMux,manager *middleware.Manager){
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
			middleware.AuthenticateJWT,	
		),
	) //advanced routing

	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProductByID),
		),
	)

	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProduct),
			middleware.AuthenticateJWT,
		),
	)

	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProduct),
			middleware.AuthenticateJWT,
		),
	)
	

}