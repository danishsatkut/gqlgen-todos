package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"

	"github.com/danishsatkut/gqlgen-todos"
	"github.com/danishsatkut/gqlgen-todos/models"
	"github.com/danishsatkut/gqlgen-todos/resolvers"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config := gqlgen_todos.Config{
		Resolvers: resolvers.NewRootResolver(models.NewDataStore()),
	}

	http.Handle("/", handler.GraphQL(gqlgen_todos.NewExecutableSchema(config)))
	http.Handle("/playground", handler.Playground("GraphQL playground", "/"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
