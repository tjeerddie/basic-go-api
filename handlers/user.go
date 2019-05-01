package handlers

import (
	"fmt"
	"log"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tjeerddie/basic-go-api/entities"
)

var (
	listQuery = "SELECT * FROM `user`"
	singleQuery = "SELECT * FROM `user` WHERE id=(?)"
	createQuery = "INSERT INTO `user` (?) VALUES (?);"
	updateQuery = "UPDATE `user` SET (?) WHERE id=(?);"
	deleteQuery = "DELETE FROM `user` WHERE id=(?);"
)


// UserList returns a list of users.
func UserList(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rows, err := db.Query(listQuery)
		if err != nil {
			log.Fatal("Database connection failed: %s", err.Error())
		}
		fmt.Printf("%+v\n", rows)

		var (
			id          int
			email       string
			password    string
			userList    []entities.User
		)

		for rows.Next() {
			err = rows.Scan(&id, &email, &password)
			userList = append(userList, entities.User{
				ID: id,
				Email: email,
				Password: password,
			})
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(userList); err != nil {
			panic(err)
		}
	}
}

// UserSingle returns a single user.
func UserSingle(db *sql.DB) (func(http.ResponseWriter, *http.Request, httprouter.Params)) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// err := db.QueryRow(singleQuery);
		// if err != nil {
		// 	log.Fatal("Database connection failed: %s", err.Error())
		// }
	}
}

// UserCreate creates a new user.
func UserCreate(db *sql.DB) (func(http.ResponseWriter, *http.Request, httprouter.Params)) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Create user!\n")
	}
}

// UserUpdate updates a existing user.
func UserUpdate(db *sql.DB) (func(http.ResponseWriter, *http.Request, httprouter.Params)) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "update, %s!\n", ps.ByName("name"))
	}
}

// UserDelete deletes a existing user.
func UserDelete(db *sql.DB) (func(http.ResponseWriter, *http.Request, httprouter.Params)) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "delete, %s!\n", ps.ByName("name"))
	}
}
