package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Sahan-g/gopher/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_"github.com/lib/pq"
)

type apiConfig struct{
	DB *database.Queries
}

func main (){
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString=="" {
		log.Fatal("PORT not found")
	}
	fmt.Println("PORT",portString)
	dbURI := os.Getenv("DB_URI")
	if dbURI=="" {
		log.Fatal("db uri not found")
	}

	conn,err:=sql.Open("postgres",dbURI)
	if err!= nil{
		log.Fatal("db connection failed",err)
	}

	apicfg := apiConfig{
		DB: database.New(conn),
	}

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
	v1Router.Post("/users",apicfg.handlerCreateUser)
	v1Router.Get("/users",apicfg.handlerGetUserByApiKey)

	router.Mount("/v1",v1Router)
	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	err=srv.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}

}