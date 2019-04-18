package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/danishsatkut/gqlgen-todos"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.GraphQL(gqlgen_todos.NewExecutableSchema(gqlgen_todos.Config{Resolvers: &gqlgen_todos.Resolver{}})))
	http.Handle("/playground", handler.Playground("GraphQL playground", "/"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
