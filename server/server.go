package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"

	"github.com/tjeerddie/basic-go-api/handlers"
	"github.com/tjeerddie/basic-go-api/repository/sqlrepo"

)

const (
	shutdownTimeout = 5 * time.Second
	serverName      = "localhost:3306"
	user            = "root"
	password        = ""
	dbName          = "basic_go_api"
)

type Server struct {
	Server     *http.Server
	Repository *sqlrepo.Repository
}

// New function creates the routes and returns the router
func New(address string) *Server {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s", user, password, serverName, dbName,
	)
	repo := sqlrepo.New(connectionString)
	router := setupRoutes(repo)
	srv := http.Server{Addr: address, Handler: router}
	return &Server{
		&srv,
		repo,
	}
}

func setupRoutes(repo *sqlrepo.Repository) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", handlers.Index)
	router.GET("/hello/:name", handlers.Hello)
	userRoutes(router, repo)
	return router
}

func userRoutes(router *httprouter.Router, repo *sqlrepo.Repository) {
	router.GET("/users", handlers.UserList(repo))
	router.GET("/users/:id", handlers.UserSingle(repo))
	router.POST("/users", handlers.UserCreate(repo))
	// router.PUT("/users/:id", handlers.UserUpdate(repo))
	// router.DELETE("/users/:id", handlers.UserDelete(repo))
}

func (s *Server) ListenAndServe() {
	s.Spawn()
	s.Block()
}

func (s *Server) Spawn() {
	go func() {
		fmt.Printf("Listening to port %+v\n", s.Server.Addr)
		if err := s.Server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
}

func (s *Server) Block() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := s.Server.Shutdown(ctx); err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		log.Println("Server stopped")
	}
}

func (s *Server) Close() {
	s.Server.Close()
	s.Repository.Close()
}
