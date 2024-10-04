package main

import {
	"fmt"
	"log"
	"io"
	"os"

	"github.com/joho/godotenv"
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLTeams?sortBy=standings&rosters=false&schedules=false&topPerformers=true&teamStats=true&teamStatsSeason=2023"
	
	apiKey := os.Getenv("RAPID_API_KEY")
	apiHost := os.Getenv("RAPID_HOST")

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", apiKey)
	req.Header.Add("x-rapidapi-host", apiHost)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}