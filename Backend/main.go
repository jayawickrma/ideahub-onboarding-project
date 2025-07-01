package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("IdeaHUb Backend Server running on port http://localhost:8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Welcome to Ideahub Backend......")
	})
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error", err)
	}
}
