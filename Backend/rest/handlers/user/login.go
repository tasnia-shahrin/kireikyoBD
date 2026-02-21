package user

import (
	"Backend/Database"
	"Backend/config"
	"Backend/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct{
	Email string `json:"email"`
	Password string `json:"password"`

}
func (h *Handler) Login(w http.ResponseWriter,r *http.Request){

	/*
	1.take body information from r.Body
	2.create an instance using Product struct with body information
	3.append the instance into productList
	*/
	var reqLogin ReqLogin //creating instance
	decoder:= json.NewDecoder(r.Body) //take body information from r.Body
	err:=decoder.Decode(&reqLogin)
	if err!=nil{
		fmt.Println(err)
		http.Error(w,"Invalid request data",http.StatusBadRequest)
		return
	}
	
	usr:=Database.Find(reqLogin.Email,reqLogin.Password)
	if usr==nil{
		http.Error(w,"Invalid credentials",http.StatusBadRequest)
		return
	}

	cnf:=config.GetConfig()
	
	accessToken,err:=util.CreateJwt(cnf.JwtSecretKey,util.Payload{
		Sub:usr.ID,
		FirstName: usr.FirstName,
		LastName: usr.LastName,
		Email: usr.Email,
	})
	if err!=nil{
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
		return
	}
	util.SendData(w,accessToken,http.StatusCreated)
	
}
