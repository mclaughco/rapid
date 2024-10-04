// The tank package interfaces with the tank api endpoints provided on RapidAPI's website
package tank

// About the Tank API
// https://rapidapi.com/tank01/api/tank01-nfl-live-in-game-real-time-statistics-nfl
//
// This API provides up-to-date data about the current NFL season
// A free account is capped at 1,000 API calls per month
// Use them wisely

// API Endpoints



import (
	// Standard pkgs
	"fmt"
	"log"
	"io"
	// "os"
	"net/http"
)

// func GetEndpoint(endpoint string, apiKey string, apiHost string) {
	
// 	// Player List
// 		// none
// 	// Weekly Schedule
// 		// week
// 		// seasonType
// 		// season
// 	// Fantasy Point Projections
// 		// week
// 		// playerID
// 		// teamID
// 		// archiveSeason
// 		// twoPointConversions
// 		// passYards
// 		// passAttempts
// 		// passTD
// 		// passCompletions
// 		// passInterceptions
// 		// pointsPerReception
// 		// carries
// 		// rushYards
// 		// rushTD
// 		// fumbles
// 		// receivingYards
// 		// receivingTD
// 		// targets
// 		// fgMade
// 		// fgMissed
// 		// xpMade
// 		// xpMissed
// 	// Roster
// 		// teamID
// 		// teamAbv
// 		// archiveDate
// 		// getStats
// 		// fantasyPoints
// 	// Average Draft Position
// 		// adpType
// 		// adpDate
// 		// pos
// 		// filterOut
// 	// Box Score
// 		// gameID
// 		// playByPlay
// 		// fantasyPoints
// 		// twoPointConversions
// 		// passYards
// 		// passAttempts
// 		// passTD
// 		// passCompletions
// 		// passInterceptions
// 		// pointsPerReception
// 		// carries
// 		// rushYards
// 		// rushTD
// 		// fumbles
// 		// receivingYards
// 		// receivingTD
// 		// targets
// 		// defTD
// 		// fgMade
// 		// fgMissed
// 	// Game Info
// 		//
// 		//
// 		//
// 	// Depth Charts
// 		// none
// 	// Betting Odds
// 		//
// 		//
// 		//
// 	// News and Headlines
// 		//
// 		//
// 		//
// 	// Scoreboard
// 		//
// 		//
// 		//
// 	// Team Schedule
// 		//
// 		//
// 		//
// 	// Daily Schedule
// 		//
// 		//
// 		//
// 	// Player Information
// 		//
// 		//
// 		//
// 	// Player Game Statistics
// 		//
// 		//
// 		//
// }

// Get NFL Teams
		// sortBy
		// rosters
		// schedules
		// topPerformers
		// teamStats
		// teamStatsSeason
func GetNFLTeams(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLTeams?sortBy=standings&rosters=false&schedules=false&topPerformers=true&teamStats=true&teamStatsSeason=2024"

	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// Get Player List
func GetPlayerList(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLPlayerList"
	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// Get Weekly Schedule
func GetWeekSchedule(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLGamesForWeek?week=1&seasonType=reg&season=2024"

	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)
	
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}


func GetFantasyProjections(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLProjections?week=season&archiveSeason=2024&twoPointConversions=2&passYards=.04&passAttempts=-.5&passTD=4&passCompletions=1&passInterceptions=-2&pointsPerReception=1&carries=.2&rushYards=.1&rushTD=6&fumbles=-2&receivingYards=.1&receivingTD=6&targets=.1&fgMade=3&fgMissed=-1&xpMade=1&xpMissed=-1"

	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// Get an NFL Team's current roster
func GetRoster(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLTeamRoster?teamID=6&teamAbv=CHI&getStats=true&fantasyPoints=true"

	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)
	
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// Get a player's Average Draft Position
func GetADP(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLADP?adpType=halfPPR"

	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)
	
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// Get the live box scores for a specific game
func GetBoxScore(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLBoxScore?gameID=20240810_CHI%40BUF&playByPlay=true&fantasyPoints=true&twoPointConversions=2&passYards=.04&passAttempts=0&passTD=4&passCompletions=0&passInterceptions=-2&pointsPerReception=.5&carries=.2&rushYards=.1&rushTD=6&fumbles=-2&receivingYards=.1&receivingTD=6&targets=0&defTD=6&fgMade=3&fgMissed=-3&xpMade=1&xpMissed=-1"

	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)
	
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// Get game information
func GetGameInfo(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLGameInfo?gameID=20240908_TEN%40CHI"

	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// Get depth charts
func GetDepthCharts(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLDepthCharts"

	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// Get betting odds
func GetBettingOdds(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLBettingOdds?gameDate=20240908&itemFormat=list&impliedTotals=true"

	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// Get news and headlines
func GetHeadlines(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLNews?fantasyNews=true&maxItems=20"

	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)
	
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// Get scoreboard
func GetScoreboard(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLScoresOnly?gameDate=20240107&topPerformers=true"

	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// Get team schedule
func GetTeamSchedule(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLTeamSchedule?teamAbv=BUF&season=2023"

	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)
	
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

// Get daily schedule
func GetDailySchedule(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLGamesForDate?gameDate=20240107"

	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)
	
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// Get player information
func GetPlayerInfo(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLPlayerInfo?playerName=keenan_a&getStats=true"

	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// Get player game statistics
func GetPlayerGameStats(apiKey string, apiHost string) {
	url := "https://tank01-nfl-live-in-game-real-time-statistics-nfl.p.rapidapi.com/getNFLGamesForPlayer?playerID=3121422&fantasyPoints=true&twoPointConversions=2&passYards=.04&passTD=4&passInterceptions=-2&pointsPerReception=1&carries=.2&rushYards=.1&rushTD=6&fumbles=-2&receivingYards=.1&receivingTD=6&targets=0&defTD=6&xpMade=1&xpMissed=-1&fgMade=3&fgMissed=-3"

	req, _ := http.NewRequest("GET", url, nil)

	req = addHeaders(req, apiKey, apiHost)
	
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error making request")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// Add the API key and API host to the headers of the request before sending it
func addHeaders(req *http.Request, apiKey string, apiHost string) *http.Request {
	req.Header.Add("x-rapidapi-key", apiKey)
	req.Header.Add("x-rapidapi-host", apiHost)
	return req
}