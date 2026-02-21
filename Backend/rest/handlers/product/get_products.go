package product

import (
	"Backend/Database"
	"Backend/util"
	"net/http"
)

//--------creating getProducts API-------
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request){

	util.SendData(w,Database.List(),200)

}
