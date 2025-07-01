// main.go
package main

import (
	"Backend/Routers"
	"Backend/utils"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.InitMongoDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default if not set
	}

	mainRouter := Routers.NewRouter()

	// Mount the router at /api/v1/
	http.Handle("/api/v1/", http.StripPrefix("/api/v1", mainRouter.Router))

	fmt.Printf("IdeaHub Backend Server running at http://localhost:%s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
