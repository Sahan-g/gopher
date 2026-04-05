package main

import (
	"database/sql"

	"github.com/Sahan-g/gopher/internal/database"
	"github.com/google/uuid"
)

type User struct{
	ID        uuid.UUID `json:"id"`
	Username  string	`json:"username"`
	Email     string	`json: "email"`
	CreatedAt sql.NullTime `json: "createdAt"`
	UpdatedAt sql.NullTime `json: updatedAt`
}

func dbUsertoUser(dbUser database.User)User{
	return User{
		Username: dbUser.Email,
	}
}