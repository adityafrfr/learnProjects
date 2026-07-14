package main

import "net/http"

var mux = http.NewServeMux()

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func init() {
	mux.HandleFunc("GET /ping", PingHandler)
	mux.HandleFunc("GET /notes", GetNoteHandler)
	mux.HandleFunc("POST /notes", CreateNotesHandler)
}
