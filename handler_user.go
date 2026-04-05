package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Sahan-g/gopher/internal/auth"
	"github.com/Sahan-g/gopher/internal/database"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

var validate = validator.New()

func (apicfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username" validate:"required,min=3,max=32"`
		Email    string `json:"email" validate:"required,email"`
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, 400, err.Error())
		return
	}

	if err := validate.Struct(params); err != nil {
		respondWithError(w, 400, err.Error())
		return
	}

	user, err := apicfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Username: params.Username,
		Email:    params.Email,
	})
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			respondWithError(w, 409, "user already exists")
			return
		}

		respondWithError(w, 500, err.Error())
		return
	}
	
	respondWithJson(w, 201, dbUsertoUser(user))
}

func (apicfg *apiConfig) handlerGetUserByApiKey(w http.ResponseWriter, r *http.Request) {
	apiKey,err := auth.GetApiKey(r.Header)
	if err != nil{
		respondWithError(w,403,err.Error())
	}
	user,err:= apicfg.DB.UserByApiKey(r.Context(),apiKey)
	if err!= nil{
		respondWithError(w,400,err.Error())
		return
	}
	respondWithJson(w,200,dbUsertoUser(user))

}