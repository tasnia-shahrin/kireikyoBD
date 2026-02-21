package user

import (
	
	"Backend/rest/middleware"
	"net/http"
	
)

func (h *Handler)RegisterRoutes(mux *http.ServeMux,manager *middleware.Manager){
	
	//-----creating user api------
	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(h.CreateUser),
		),
	)
	//----login api----
	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(h.Login),
		),
	)

}