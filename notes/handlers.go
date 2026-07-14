package main

import (
	"net/http"
)


func pingHandler(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	response := "all ok"
	w.Write([]byte(response))
}

// 		response := map[string]string{
// 		"status": "all ok",
// 	}

// 	if err := json.NewEncoder(w).Encode(response); err != nil {
// 		log.Printf("failed to encode response: %v", err)
// 	}
// }

