package main

import (
	"io"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	// Request environment variables
	accessToken    string
	clanTag        string
	encodedClanTag string
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Fehler beim Laden der .env-Datei")
	}

	gin.SetMode(gin.ReleaseMode)

	logFile, err := os.OpenFile("gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Fehler beim Öffnen der Log-Datei:", err)
	}
	defer logFile.Close()

	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	router := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Request environment variables
	accessToken = os.Getenv("ACCESS_TOKEN")
	clanTag = os.Getenv("CLAN_TAG")
	encodedClanTag = url.QueryEscape(clanTag)

	// Start data collector
	go func() {
		for {
			dataCollector()
			time.Sleep(time.Minute)
		}
	}()

	router.GET("/api/clan", getClanHandler)
	router.GET("/api/members", getMembersHandler)
	router.GET("/api/currentriverrace", getCurrentRiverRaceHandler)
	router.GET("/api/riverracelog", getRiverRaceLogHandler)
	router.GET("/database/person", getPerson)
	router.GET("/database/clan", getClan)
	router.POST("/database/person", createPersonManuell)

	// Enable CORS
	router.Use(cors.Default())

	log.Printf("Server läuft auf Port %s", port)
	log.Fatal(router.Run(":" + port))
}
