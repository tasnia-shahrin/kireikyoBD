package handlers

import (
	"Backend/Database"
	"Backend/util"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter,r *http.Request){
	productID:=r.PathValue("id")

	pId,err:=strconv.Atoi(productID)
	if err!=nil{
		http.Error(w,"please give me a valid product id",400)
		return
	}

	Database.Delete(pId)

	util.SendData(w,"Successfully deleted product",201)
}