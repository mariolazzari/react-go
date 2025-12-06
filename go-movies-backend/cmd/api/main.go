package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// app config
	var app application

	// command ling args

	// connct db

	app.Domain = "example.com"

	log.Println("Starting application on port", port)
	// start web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
