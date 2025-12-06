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

```
