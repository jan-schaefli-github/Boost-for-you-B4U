package main

import (
	//Local Packages Endpoint Import
	"b4u/backend/endpoints/v1/ep_api/aep_clan"
	"b4u/backend/endpoints/v1/ep_database/dbep_clan"
	"b4u/backend/endpoints/v1/ep_database/dbep_person"
	"b4u/backend/logger"
	"b4u/backend/routine/v1/rt_main"
	"b4u/backend/tools"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Create logs directory if it doesn't exist
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0755)
		if err != nil { // If there was an error while creating the directory
			logger.LogMessage("Logs", "Error while creating logs directory: "+err.Error())
			return
		}
	}

	tools.LoadDotEnv()
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Create gin log file
	logFileGin, err := os.OpenFile(filepath.Join("logs", "gin.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // Open gin log file in append mode
	if err != nil {                                                                                             // If there was an error while creating the gin log file
		logger.LogMessage("Gin", "Error while creating gin log file: "+err.Error())
	}

	// Close gin log file
	defer func(logFileGin *os.File) {
		err := logFileGin.Close()
		if err != nil {
			logger.LogMessage("Gin", "Error while closing gin log file: "+err.Error())
		}
	}(logFileGin)

	// Log to file and console
	gin.DefaultWriter = io.MultiWriter(logFileGin, os.Stdout)

	// Create router
	router := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	//Start routine
	go func() {
		for {
			go rt_main.Routine()
			time.Sleep(time.Hour)
		}
	}()

	// Routes

	router.GET("/api/clan", aep_clan.GetClan)
	router.GET("/api/clan/members", aep_clan.GetMembers)
	router.GET("/api/clan/currentriverrace", aep_clan.GetCurrentRiverRace)
	router.GET("/api/clan/riverracelog", aep_clan.GetRiverRaceLog)
	router.GET("/api/clan/members/leaderboard", aep_clan.GetClanMemberLeaderboard)
	router.GET("/api/clan/locations", aep_clan.GetLocations)
	router.GET("/api/clan/leaderboard", aep_clan.GetClanRankingByLocation)
	router.GET("/database/person", dbep_person.GetPerson)
	router.GET("/database/clan", dbep_clan.GetClan)

	// Enable CORS
	router.Use(cors.Default())

	// Start server
	log.Printf("Server läuft auf Port %s", port)
	log.Fatal(router.Run(":" + port))
}
