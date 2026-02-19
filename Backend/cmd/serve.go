package cmd

import (
	"Backend/config"
	"Backend/rest"
)

func Serve(){

	cnf := config.GetConfig()

	rest.Start(cnf)

}