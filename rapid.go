package main

import (
	// Standard pkgs
	// "fmt"
	"log"
	// "io"
	"os"
	// "net/http"

	// My pkgs
	"github.com/mclaughco/rapid/tank"

	// Other people's pkgs
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file in the directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	apiKey := os.Getenv("RAPID_API_KEY")
	apiHost := os.Getenv("RAPID_HOST")

	tank.GetNFLTeams(apiKey, apiHost)

}