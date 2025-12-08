package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DB     *sql.DB
	DSN    string
	Domain string
}

func main() {
	// app config
	var app application
	// command ling args
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=mario password=password dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// connct db
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = conn

	app.Domain = "example.com"

	log.Println("Starting application on port", port)
	// start web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
