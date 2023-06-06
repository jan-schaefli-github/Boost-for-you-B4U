package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	// Database environment variables
	username string
	password string
	host     string
	port     int
)

// Connect to MySQL database
func connectToDatabase() (*sql.DB, error) {
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Get Clan from Database
func getClan(c *gin.Context) {
	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}
	defer db.Close()

	var clan Clan
	err = db.QueryRow("SELECT * FROM clan WHERE tag = '#P9UVQCJV'").Scan(&clan.Tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	clanJSON, err := json.Marshal(clan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(clanJSON))
}

// Get Person from Database
func getPerson(c *gin.Context) {
	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}
	defer db.Close()

	var person Person
	err = db.QueryRow("SELECT * FROM person WHERE tag = '#2Y9VQVJ8'").Scan(&person.Tag, &person.WholeFame, &person.ClanStatus, &person.JoinDate, &person.FkClan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	personJSON, err := json.Marshal(person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(personJSON))
}
