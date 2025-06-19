package main

import (
	"log"
	"net/http"

	db "smartpill/connection"
)

func main() {
	db.ConnectDatabase()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
