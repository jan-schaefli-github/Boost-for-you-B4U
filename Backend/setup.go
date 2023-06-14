package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"os"
	"path/filepath"
)

func setup() {
	// Create logs directory if it doesn't exist
	if _, err := os.Stat("ApiLogs"); os.IsNotExist(err) {
		err := os.Mkdir("ApiLogs", 0755)
		if err != nil { // If there was an error while creating the directory
			logMessage("Logs", "Error while creating logs directory: "+err.Error())
			return
		}
	}

	// Load environment variables out of .env file
	err := godotenv.Load()
	if err != nil {
		logMessage("Environment", "Error while loading environment variables: "+err.Error())
	}

	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Create gin log file if not existent
	logFileGin, err := os.OpenFile(filepath.Join("ApiLogs", "gin.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logMessage("Gin", "Error while creating gin log file: "+err.Error())
	}

	// Close gin log file connection
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
		port = "3000" // Default port if not specified
	}
}
