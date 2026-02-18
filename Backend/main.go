package main

import (
	"Backend/cmd"
	
)

//---handling CORS----
// func handleCORS(w http.ResponseWriter){
// 	w.Header().Set("Access-Control-Allow-Origin","*")
// 	w.Header().Set("Access-Control-Allow-Methods",",GET,POST,PUT,PATCH,DELETE,OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers","Content-Type")
// 	w.Header().Set("Content-Type","application/json")
// }
// //----handling preflight requests-----
// func handlePreflightReq(w http.ResponseWriter,r *http.Request){
// 	if r.Method=="OPTIONS"{
// 		w.WriteHeader(200)
// 		return
// 	}
// }


func main(){
	
	cmd.Serve()
}

