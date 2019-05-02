package service

import (
	"log"
	"fmt"
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"

	"github.com/tjeerddie/basic-go-api/handlers"
)


type Server struct {
	SRV *http.Server
	DB *sql.DB
}

// New function creates the routes and returns the router
func New(address string) *Server {
	db := setupDB()
	router := setupRoutes(db)
	srv := http.Server{Addr: address, Handler: router}
	return &Server{
		&srv,
		db,
	}
}

func setupDB() *sql.DB {
	serverName := "localhost:3306"
	user := "root"
	password := ""
	dbName := "basic_go_api"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, serverName, dbName)

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
	router.GET("/users/:id", handlers.UserSingle(db))
	router.POST("/users", handlers.UserCreate(db))
	router.PUT("/users/:id", handlers.UserUpdate(db))
	router.DELETE("/users/:id", handlers.UserDelete(db))
}
