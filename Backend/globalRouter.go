package main

import(
	
	"net/http"
	
)

//------global router: handle cors and preflight request--------
func globalRouter(mux *http.ServeMux) http.Handler{
	handleAllReq:=func(w http.ResponseWriter,r *http.Request){
		//handling CORS
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Header().Set("Access-Control-Allow-Methods",",GET,POST,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers","Content-Type")
		w.Header().Set("Content-Type","application/json")
		//Handling preflight req
		if r.Method=="OPTIONS"{
			w.WriteHeader(200)
			return
		}else{
			mux.ServeHTTP(w,r)
		}
	}
	return http.HandlerFunc(handleAllReq)
}
