package main

import (
	"api/internal/api"
	"log"
	"net/htpp"
	"os"
	//"github.com/gin-gonic/gin"
)

func login(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Login")
}

func main() {
	api.Initialize()
	
	mux.HandleFunc("/", home)
	
	mux := http.NewServeMux()

	mux.HandleFunc("/handlers/auth/login.go", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("/login response"))
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			htpp.NotFound(w, r)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "6060"
	}

	log.Printf("Running in http://localhost:%d", port)
	if err := api.Router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}