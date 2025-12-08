# Working with React and Go (Golang)

## React

## Go

```sh
go mod init backend
```

### Starting backend API

```go
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
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Routes and handlers

```go
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
	http.HandleFunc("/", Hello)
	// start web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Intalling chi package

```sh
go get -u github.com/go-chi/chi/v5
```

```go
package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *application) routes() http.Handler {
	// create new router
	mux := chi.NewRouter()
	// middlewares
	mux.Use(middleware.Recoverer)

	return mux
}
```

### Adding route handler

```go
package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *application) routes() http.Handler {
	// create new router
	mux := chi.NewRouter()
	// middlewares
	mux.Use(middleware.Recoverer)
	// routes
	mux.Get("/", app.Home)

	return mux
}
```

### Returning JSON from API

```go
	var payload = struct {
		Status string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status: "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
```

### Returning list of movies in JSON

```go
func (app *application) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	rd, _ := time.Parse("2020-01-01", "2020-03-28")
	movie := models.Movie{
		ID:          1,
		Title:       "titolo",
		ReleaseDate: rd,
		MPAARating:  "R",
		RunTime:     111,
		Description: "movie description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now()}

	movies = append(movies, movie)
	movies = append(movies, movie)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
```

### Enabling CORS

```go
package main

import "net/http"

func (app *application) enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://*")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, X-CSRF-Token, Authorization")
			return
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
```

## Connecting to Postgres

### Docker setup

```yaml
version: "3"

services:
  postgres:
    image: "postgres:14.5"
    restart: always
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: movies
    logging:
      options:
        max-size: 10m
        max-file: "3"
    ports:
      - "5432:5432"
    volumes:
      - ./postgres-data:/var/lib/postgresql/data
      - ./sql/create_tables.sql:/docker-entrypoint-initdb.d/create_tables.sql
```

### Connecting API to Postgres

[pgx](https://github.com/jackc/pgx)

```sh
go get github.com/jackc/pgx/v5
```

```go

```
