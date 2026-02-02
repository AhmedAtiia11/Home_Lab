package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"msg": "healthy",
	})
	fmt.Println("Health check endpoint hit")
}

func main() {
	http.HandleFunc("/health", createUser)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
