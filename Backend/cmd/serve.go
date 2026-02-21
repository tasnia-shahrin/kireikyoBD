package cmd

import (
	"Backend/config"
	"Backend/rest"
	"Backend/rest/handlers/product"
	"Backend/rest/handlers/user"
)

func Serve(){

	cnf := config.GetConfig()

	productHandler:=product.NewHandler()
	userHandler:=user.NewHandler()

	server:=rest.NewServer(productHandler,userHandler)
	server.Start(cnf)
}