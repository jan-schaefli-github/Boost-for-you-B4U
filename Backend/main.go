package main

import (
	"github.com/gin-contrib/cors"
	"log"
	"net/url"
	"os"
	"time"
)

var (
	//Define Environment Variables
	accessToken    string
	clanTag        string
	encodedClanTag string
)

func main() {
	// Call Setup function to initiate Programm
	setup()

	// Request environment variables
	accessToken = os.Getenv("ACCESS_TOKEN")
	clanTag = os.Getenv("CLAN_TAG")
	encodedClanTag = url.QueryEscape(clanTag)

	// Start routine
	go func() {
		for {
			go dataCollector(getClanTags())
			time.Sleep(time.Hour)
		}
	}()

	// Routes
	router.GET("/api/clan", getClanHandler)
	router.GET("/api/members", getMembersHandler)
	router.GET("/api/currentriverrace", getCurrentRiverRaceHandler)
	router.GET("/api/riverracelog", getRiverRaceLogHandler)
	router.GET("/api/members/leaderboard", getClanMemberLeaderboardHandler)
	router.GET("/api/locations", getLocationsHandler)
	router.GET("/api/clan/leaderboard", getRankingByLocationHandler)
	router.GET("/database/person", getPerson)
	router.GET("/database/clan", getClan)

	// Enable CORS
	router.Use(cors.Default())

	// Start server
	log.Printf("Server läuft auf Port %s", port)
	log.Fatal(router.Run(":" + port))
}
