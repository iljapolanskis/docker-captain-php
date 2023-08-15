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

	fmt.Println("Server started on port 80")
	http.ListenAndServe(":80", nil)
}

