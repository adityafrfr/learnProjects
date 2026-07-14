package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("starting on 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("something wrong")
	}	
}
