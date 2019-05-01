package service

import (
	"log"
	"fmt"
	"database/sql"
	"github.com/julienschmidt/httprouter"

	_ "github.com/go-sql-driver/mysql"

	"github.com/tjeerddie/basic-go-api/handlers"
)


type Server struct {
	Router *httprouter.Router
	DB *sql.DB
}

// New function creates the routes and returns the router
func New() *Server {
	var srv Server
	srv.DB = setupDB()
	srv.Router = setupRoutes(srv.DB)
	return &srv
}

func setupDB() *sql.DB {
	serverName := "localhost:3306"
	user := "root"
	password := ""
	dbName := "basic_go_api"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	return db
}

func setupRoutes(db *sql.DB) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", handlers.Index)
	router.GET("/hello/:name", handlers.Hello)
	userRoutes(router, db)
	return router
}

func userRoutes(router *httprouter.Router, db *sql.DB) {
	router.GET("/users", handlers.UserList(db))
	router.GET("/users/:name", handlers.UserSingle(db))
	router.POST("/users", handlers.UserCreate(db))
	router.PUT("/users/:name", handlers.UserUpdate(db))
	router.DELETE("/users/:name", handlers.UserDelete(db))
}
