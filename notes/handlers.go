package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	response := "all ok"
	w.Write([]byte(response))
}

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {

	notes := publicNotes.Read()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)


}


func CreateNotesHandler(w http.ResponseWriter, r *http.Request) {

	var req CreateNoteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil	{
		log.Printf("cannot create note with %v %q", r.Body, err)

		http.Error(w, "cant create the note", http.StatusBadRequest)
		return
	}

	newNote := publicNotes.Create(req)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newNote)
}

// 		response := map[string]string{
// 		"status": "all ok",
// 	}

// 	if err := json.NewEncoder(w).Encode(response); err != nil {
// 		log.Printf("failed to encode response: %v", err)
// 	}
// }
