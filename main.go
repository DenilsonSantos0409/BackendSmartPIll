package main

import (
	"log"
	"net/http"

	db "smartpill/connection" 
	"smartpill/schema"
	"github.com/graphql-go/handler"
)

func main() {
	db.ConnectDatabase()
	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)

	log.Println("🚀 GraphQL siap di http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
