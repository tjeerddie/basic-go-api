package main

import (
	"fmt"

	"github.com/tjeerddie/basic-go-api/config"
	serverService "github.com/tjeerddie/basic-go-api/server"
)

var configFile = ".env"
var defaultPort = "8000"

func main() {
	config.ReadDotEnv(configFile)
	address := fmt.Sprintf(":%s", config.Getenv("PORT", defaultPort))
	server := serverService.New(address)
	defer server.Close()
	server.ListenAndServe()
}
