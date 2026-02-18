package cmd

import (
	"Backend/config"
	"Backend/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Serve(){

	cnf := config.GetConfig()

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
	initRoutes(mux,manager)
	
	addr:= ":" + strconv.Itoa(cnf.HttpPort)
		fmt.Println("Server running on port ",addr)
	err:=http.ListenAndServe(addr,wrappedMux)
	if err!=nil{
		fmt.Println("Error occurs while starting the server",err)
		os.Exit(1)
	}

}