package main

import (
	"net/http"
	"log"

	"github.com/tjeerddie/basic-go-api/service"
)

func main() {
	server := service.New()
	log.Fatal(http.ListenAndServe(":8080", server.Router))
}
