package main

import (
	"Backend/utils"
	"fmt"
	"net/http"
)

func main() {
	utils.InitMongoDB()
	fmt.Println("IdeaHUb Backend Server running on port http://localhost:8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Welcome to IdeaHub Backend......")
	})
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error", err)
	}

}
