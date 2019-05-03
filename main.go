package main

import (
	"fmt"

	"github.com/tjeerddie/basic-go-api/config"
	"github.com/tjeerddie/basic-go-api/service"
)

var configFile = ".env"
var defaultPort = "8000"

func main() {
	config.ReadDotEnv(configFile)
	address := fmt.Sprintf(":%s", config.Getenv("PORT", defaultPort))
	server := service.New(address)
	defer server.Close()
	server.ListenAndServe()
}
