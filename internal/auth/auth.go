package auth

import (
	"errors"
	"net/http"
	"strings"
)

//extractAPIkey from header
//Authorization: APIKEY {ACTUAL_API_KEY}
func GetApiKey(header http.Header)(string,error){
	val:=header.Get("Authorization") 
	if val==""{
		return "",errors.New("Authorization Header not found")
	}
	vals:=strings.Split(val," ")
	if len(vals)!=2 {
		return  "",errors.New("Malformed Auth Header")
	}
	if vals[0] != "APIKEY"{
		return "", errors.New("Invalid API Key")
	}
	return vals[1],nil
}