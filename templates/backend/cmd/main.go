package main

import (
	"log"
	"net/http"
	"os"

	backend "github.com/gonzabosio/vgstack-cli/templates/backend/api"

	"github.com/joho/godotenv"
)

func main() {
	// UNCOMMENT .ENV FROM THE GITIGNORE FILE

	if err := godotenv.Load("../backend/.env"); err != nil {
		log.Printf("environment variables loading error: %v\n", err)
	}
	r, err := backend.NewRouter()
	if err != nil {
		log.Fatalf("Unable to setup backend server: %v", err)
	}
	if err := http.ListenAndServe(":"+os.Getenv("BACKEND_PORT"), r); err != nil {
		log.Fatalf("Unable to start backend server on port %s: %v", os.Getenv("BACKEND_PORT"), err)
	}
}
