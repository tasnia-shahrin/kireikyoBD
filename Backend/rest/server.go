package rest

import (
	"Backend/config"
	"Backend/rest/handlers/user"

	"Backend/rest/handlers/product"
	"Backend/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct{
	//keeping productHandler,userHandler as dependency
	productHandler *product.Handler
	userHandler *user.Handler
}
//injecting dependencies inside NewServer()
func NewServer(
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server{
	return &Server{
		productHandler: productHandler,
		userHandler: userHandler,
	}
}
func (server *Server) Start(cnf config.Config){

	
	//-----Manager handles Middleware----
	manager:=middleware.NewManager()
	
	//globalrouter:= middleware.CorsWithPreflight(mux)

	//Preflight(CORS(logger(mux)))
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)
	//-----creating router-----
	mux:=http.NewServeMux()
	wrappedMux:=manager.WrapMux(mux)
	
	server.productHandler.RegisterRoutes(mux,manager)
	server.userHandler.RegisterRoutes(mux,manager)
	
	addr:= ":" + strconv.Itoa(cnf.HttpPort)
		fmt.Println("Server running on port ",addr)
	err:=http.ListenAndServe(addr,wrappedMux)
	if err!=nil{
		fmt.Println("Error occurs while starting the server",err)
		os.Exit(1)
	}
}