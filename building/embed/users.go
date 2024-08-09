package main

import (
	"database/sql"
	_ "embed"
)

//
//go:embed user.sql
var userSQL string

// User is a user in the system.
type User struct {
	ID    string //UUID
	Name  string
	Email string
}

// UserByID returns a user by ID.
func UserByID(db *sql.DB, id string) (User, error) {
	row := db.QueryRow(userSQL, sql.Named("id", id))
	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Email)
	return u, err
}
