package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter,code int,payload interface{} ){
	data,err := json.Marshal(payload)
	if err!= nil{ 
		w.WriteHeader(500)
		log.Printf("failed to marshal JSON reponse %v",payload)
		return
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter,code int,message string){
	if code >499 {
		log.Println("Responding with 5XX error",message)
	}
	type errorResponse struct{
		Error string `json:"error"`
	}
	respondWithJson(w,code,errorResponse{
		Error: message,
	})
	
}