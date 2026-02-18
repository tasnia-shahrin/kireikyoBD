package cmd

import (
	
	
	"Backend/middleware"
	"fmt"
	"net/http"
)

func Serve(){

	
	manager:=middleware.NewManager()
	
	//-----creating router-----
	mux:=http.NewServeMux()


	
	//globalrouter:= middleware.CorsWithPreflight(mux)

	//CorsWithPreflight(logger(mux))
	manager.Use(
		middleware.CorsWithPreflight,
		middleware.Logger,
	)
	wrappedMux:=manager.WrapMux(mux)
	initRoutes(mux,manager)
	fmt.Println("Server running on port: 8080")
	
	err:=http.ListenAndServe(":8080",wrappedMux)
	if err!=nil{
		fmt.Println("Error occurs while starting the server",err)
	}

}