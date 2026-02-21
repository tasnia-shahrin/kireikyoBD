package product

import (
	"Backend/Database"
	"Backend/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductByID(w http.ResponseWriter,r *http.Request){
	productID:= r.PathValue("id")
	pId,err:= strconv.Atoi(productID)
	if err!=nil{
		http.Error(w,"Please give me a valid product id",400)
		return
	}

	product:=Database.Get(pId)

	if product==nil{
		util.SendError(w,404,"Product not found")
		return
	}
	// for _,product:= range Database.ProductList{
	// 	if product.ID == pId{
	// 		util.SendData(w,product,200)
	// 		return
	// 	}
		
	// }
	util.SendData(w,product,200)
	

}