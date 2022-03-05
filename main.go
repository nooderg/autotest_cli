package main

import (
	"autotest-cli/authentication"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", authentication.Login)
	log.Println("Ecoute en http://127.0.0.1:8080")
	http.ListenAndServe(":8080", mux)
}
