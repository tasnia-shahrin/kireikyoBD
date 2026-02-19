package handlers

import (
	"Backend/Database"
	"Backend/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter,r *http.Request){

	/*
	1.take body information from r.Body
	2.create an instance using Product struct with body information
	3.append the instance into productList
	*/
	var newUser Database.User //creating instance
	decoder:= json.NewDecoder(r.Body) //take body information from r.Body
	err:=decoder.Decode(&newUser)
	if err!=nil{
		fmt.Println(err)
		http.Error(w,"Invalid request data",http.StatusBadRequest)
		return
	}
	//append the instance into user list
	createdUser:= newUser.Store()
	
	util.SendData(w,createdUser,http.StatusCreated)
	
}
