package main

import (
	"account-service/config"
	_ "account-service/docs"
	"account-service/router"
	"log"
	"net/http"
	"time"
)

// @title Swagger Account API
// @version 1.0
// @description This is a sample server Account API.
// @host localhost:18888
// @BasePath /api/v1
func main() {

	server := http.Server{
		Addr:         "localhost:" + config.GetAppPort(),
		Handler:      router.Router(),
		ReadTimeout:  6000 * time.Second,
		WriteTimeout: 6000 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
