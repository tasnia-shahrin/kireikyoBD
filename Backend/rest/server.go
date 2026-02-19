package rest

import (
	"Backend/config"
	
	"Backend/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config){

	
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
	InitRoutes(mux,manager)
	
	addr:= ":" + strconv.Itoa(cnf.HttpPort)
		fmt.Println("Server running on port ",addr)
	err:=http.ListenAndServe(addr,wrappedMux)
	if err!=nil{
		fmt.Println("Error occurs while starting the server",err)
		os.Exit(1)
	}
}