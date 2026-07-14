package main

import "net/http"

var mux = http.NewServeMux()

func init() {
	mux.HandleFunc("GET /ping", pingHandler)
}
