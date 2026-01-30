package main

import (
	"encoding/json"
	"net/http"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"msg": "healthy",
	})
}

func main() {

	// c := cors.New(cors.Options{
	//     AllowedOrigins:   []string{"https://example.com"},
	//     AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	//     AllowedHeaders:   []string{"Content-Type", "Authorization"},
	//     AllowCredentials: true,
	// })
	http.HandleFunc("/health", createUser)
	http.ListenAndServe("localhost:8080", nil)
}
