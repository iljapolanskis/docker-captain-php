package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"message": "Hello, World!",
		}
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

