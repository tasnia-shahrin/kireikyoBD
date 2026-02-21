package product

import (
	"Backend/rest/middleware"
	"net/http"
	
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux,manager *middleware.Manager){
	//---routes for getting products---
	mux.Handle(
		"GET /products",
		
		manager.With(
			http.HandlerFunc(h.GetProducts),
		),
	)

	//----routes for creating products-----
	// mux.HandleFunc("/create-products",createProducts)-->previous routing
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProducts),	
			middleware.AuthenticateJWT,	
		),
	) //advanced routing

	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProductByID),
		),
	)

	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProduct),
			middleware.AuthenticateJWT,
		),
	)

	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			middleware.AuthenticateJWT,
		),
	)
	

}