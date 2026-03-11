package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "80"

type Config struct{}

func main() {

	// this will be config args for the server
	app := Config{}

	log.Printf("Starting broker service on port %s\n", PORT)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
