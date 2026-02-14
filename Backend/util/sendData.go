package util

import(
	
	"net/http"
	"encoding/json"
)

//Sending data
func SendData(w http.ResponseWriter,data interface{},statusCode int){
	w.WriteHeader(statusCode)
	encoder:=json.NewEncoder(w)
	encoder.Encode(data)
}
