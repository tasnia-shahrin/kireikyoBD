package middleware

import (
	"Backend/config"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

func AuthenticateJWT(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*For authorization:
	1.parse jwt
	2.parse header & payload or claims
	3.using hmac-sha256 algorithm->create hash hmac(header,payload,secret key)
	4.parse signature part from the jwt
	5.if signature and hash is same=>forward to create products
	6.otherwise send 401 status code with unauthorized
	*/

	header:=r.Header.Get("Authorization")
	if header==""{
		http.Error(w,"Unauthorized",http.StatusUnauthorized)
		return
	}
	headerArr:=strings.Split(header," ")
	if len(headerArr)!=2{
		http.Error(w,"Unauthorized",http.StatusUnauthorized)
		return
	}
	accessToken:= headerArr[1]
	tokenParts:=strings.Split(accessToken,".")
	if len(tokenParts)!=3{
		http.Error(w,"Unauthorized",http.StatusUnauthorized)
		return
	}
	jwtHeader:=tokenParts[0]
	jwtPayload:=tokenParts[1]
	signature:=tokenParts[2]

	//connecting base64 header and payload:
	message:= jwtHeader+"."+jwtPayload
	byteArrMessage:=[]byte(message)
	cnf:=config.GetConfig()
	byteArrSecret:=[]byte(cnf.JwtSecretKey)//secret key
	
	//creating signature:
	h:=hmac.New(sha256.New,byteArrSecret)
	h.Write(byteArrMessage)

	hash:=h.Sum(nil) //it gives hash as output
	newSignature:=base64UrlEncode(hash)

	if newSignature != signature{
		http.Error(w,"Unauthorized.HACKER!!",http.StatusUnauthorized)
		return
	}
	next.ServeHTTP(w,r)
	})
}
func base64UrlEncode(data []byte) string{
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}