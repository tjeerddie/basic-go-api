package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/tjeerddie/basic-go-api/entities"
	"github.com/tjeerddie/basic-go-api/repository"
	resphdr "github.com/tjeerddie/basic-go-api/handlers/responsehandler"
)

// UserList returns a list of users.
func UserList(repo repository.Repository) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		users, err := repo.Users()
		if err != nil {
			resphdr.WriteErrorResponse(
				w, http.StatusInternalServerError, fmt.Sprintf("Internal server error: %s", err.Error()),
			)
			return
		}

		resphdr.WriteOKResponse(w, users)
	}
}

// UserSingle returns a single user.
func UserSingle(repo repository.Repository) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			resphdr.WriteErrorResponse(
				w, http.StatusInternalServerError, fmt.Sprintf("Internal server error: %s", err.Error()),
			)
			return
		}

		user, err := repo.UserSingle(id)
		if err != nil {
			resphdr.WriteErrorResponse(
				w, http.StatusInternalServerError, err.Error(),
			)
			return
		}

		fmt.Printf("%+v\n", user)
		resphdr.WriteOKResponse(w, user)
	}
}

// UserCreate creates a new user.
func UserCreate(repo repository.Repository) httprouter.Handle {
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

		err = repo.UserStore(&user)
		if err != nil {
			resphdr.WriteErrorResponse(
				w, http.StatusInternalServerError, err.Error(),
			)
			return
		}

		resphdr.WriteOKResponse(w, user)
	}
}

// UserUpdate updates a existing user.
func UserUpdate(repo repository.Repository) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "update, %s!\n", ps.ByName("name"))
	}
}

// UserDelete deletes a existing user.
func UserDelete(repo repository.Repository) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "delete, %s!\n", ps.ByName("name"))
	}
}
