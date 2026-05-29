package main

import (
	"api/internal/api"
	"log"
	"os"

)

func main() {
	api.Initialize()

	port := os.Getenv("PORT")
	if port == "" {
		port = "6060"
	}

	log.Printf("Running in http://localhost:%s", port)
	if err := api.Router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}