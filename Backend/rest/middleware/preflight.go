package middleware

import(
	
	"net/http"
	
)

//------global router: handle cors and preflight request--------
func Preflight(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		
		//Handling preflight req
		if r.Method=="OPTIONS"{
			w.WriteHeader(200)
			return
		}
		next.ServeHTTP(w,r)
		
	})
	
}
