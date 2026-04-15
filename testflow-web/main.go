package main

import (
	"log"
	"net/http"

	"testflow-web/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/health", handlers.Health)

	log.Println("testflow-web listening on :9090")
	if err := http.ListenAndServe(":9090", mux); err != nil {
		log.Fatal(err)
	}
}
