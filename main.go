package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main (){
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString=="" {
		log.Fatal("PORT not found")
	}
	fmt.Println("PORT",portString)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*","https://*"}, // 
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router :=  chi.NewRouter()
	v1Router.Get("/healthz",handlerRediness)
	v1Router.Get("/error",handlerError)

	router.Mount("/v1",v1Router)
	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	err:=srv.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}

}