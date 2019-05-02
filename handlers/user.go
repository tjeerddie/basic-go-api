package handlers

import (
	"fmt"
	"log"
	"database/sql"
	"encoding/json"
	"net/http"
	"io/ioutil"

	"github.com/julienschmidt/httprouter"
	"github.com/tjeerddie/basic-go-api/entities"
	resphdr "github.com/tjeerddie/basic-go-api/handlers/responsehandler"
)

var (
	listQuery = "SELECT * FROM `user`"
	singleQuery = "SELECT * FROM `user` WHERE id=?"
	createQuery = "INSERT INTO `user` (email, password) VALUES (?) RETURNING id"
	updateQuery = "UPDATE `user` SET (?) WHERE id=?"
	deleteQuery = "DELETE FROM `user` WHERE id=?"
)


// UserList returns a list of users.
func UserList(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rows, err := db.Query(listQuery)
		if err != nil {
			log.Fatal("Database connection failed: %s", err.Error())
			return;
		}

		var (
			id          int
			email       string
			password    string
			userList    []entities.User
		)

		for rows.Next() {
			err = rows.Scan(&id, &email, &password)
			userList = append(userList, entities.User{
				Id: id,
				Email: email,
				Password: password,
			})
		}

		resphdr.WriteOKResponse(w, userList)
	}
}

// UserSingle returns a single user.
func UserSingle(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var (
			findId string = ps.ByName("id")
			id int
			email string
			password string
			user entities.User
		)
		row := db.QueryRow(singleQuery, findId)

		if err := row.Scan(&id, &email, &password); err != nil {
			if err == sql.ErrNoRows {
				resphdr.WriteErrorResponse(
					w, http.StatusNotFound, "User not found: " + findId,
				)
				return
			} else {
				log.Fatal("Database connection failed: %s", err.Error())
				return
			}
		}
		user = entities.User{
			Id: id,
			Email: email,
			Password: password,
		}

		fmt.Printf("%+v\n", user)
		resphdr.WriteOKResponse(w, user)
	}
}

// UserCreate creates a new user.
func UserCreate(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			resphdr.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var user entities.User
		if err := json.Unmarshal(body, &user); err != nil {
			resphdr.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		row := db.QueryRow(createQuery, user.Email, user.Password)
		if err := row.Scan(&user.Id); err != nil {
			resphdr.WriteErrorResponse(
				w, http.StatusInternalServerError, "Internal Server Error",
			)
			return
		}

		resphdr.WriteOKResponse(w, user)
	}
}

// UserUpdate updates a existing user.
func UserUpdate(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "update, %s!\n", ps.ByName("name"))
	}
}

// UserDelete deletes a existing user.
func UserDelete(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "delete, %s!\n", ps.ByName("name"))
	}
}
