package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", pingHandler)

	print("starting on 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("something wrong")
	}

	
}
