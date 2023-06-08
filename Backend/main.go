package main

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	// Request environment variables
	accessToken    string
	clanTag        string
	encodedClanTag string
)

func main() {

	// Create logs directory if it doesn't exist
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0755)
		if err != nil {
			logMessage("Logs", "Error while creating logs directory: "+err.Error())
			return
		}
	}

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		logMessage("Environment", "Error while loading environment variables: "+err.Error())
	}

	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Create gin log file
	logFileGin, err := os.OpenFile(filepath.Join("logs", "gin.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logMessage("Gin", "Error while creating gin log file: "+err.Error())
	}
	defer func(logFileGin *os.File) {
		err := logFileGin.Close()
		if err != nil {
			logMessage("Gin", "Error while closing gin log file: "+err.Error())
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

	// Request environment variables
	accessToken = os.Getenv("ACCESS_TOKEN")
	clanTag = os.Getenv("CLAN_TAG")
	encodedClanTag = url.QueryEscape(clanTag)

	// Start routine
	go func() {
		for {
			dataCollector()
			time.Sleep(time.Minute)
		}
	}()

	// Routes
	router.GET("/api/clan", getClanHandler)
	router.GET("/api/members", getMembersHandler)
	router.GET("/api/currentriverrace", getCurrentRiverRaceHandler)
	router.GET("/api/riverracelog", getRiverRaceLogHandler)
	router.GET("/database/person", getPerson)
	router.GET("/database/clan", getClan)

	// Enable CORS
	router.Use(cors.Default())

	// Start server
	log.Printf("Server läuft auf Port %s", port)
	log.Fatal(router.Run(":" + port))
}

// Log message to file
func logMessage(logType string, message string) {

	// Open log file in append mode
	logFile, err := os.OpenFile(filepath.Join("logs", "logfile.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error while creating log file: " + err.Error())
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Println("Error while closing log file: " + err.Error())
		}
	}(logFile)

	// Log to file
	logger := log.New(logFile, "", log.Ldate|log.Ltime)
	logString := fmt.Sprintf("[%s] %s", logType, message)
	logger.Println(logString)
}
