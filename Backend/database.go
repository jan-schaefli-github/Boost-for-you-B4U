package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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

////--------------------------------------------- Get ---------------------------------------------

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

//--------------------------------------------- Create ---------------------------------------------

// Create Person in Database
func createPerson(c *gin.Context) {
	log.Println("Received request to create a person")
	// Set CORS headers
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight OPTIONS request
	if c.Request.Method == "OPTIONS" {
		c.Status(http.StatusOK)
		return
	}

	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}
	defer db.Close()

	// Parse request body
	// Parse request body
	var person Person
	if err := c.ShouldBindJSON(&person); err != nil {
		log.Println("Error while parsing JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	log.Printf("Received person data: %+v", person)

	// Insert person into the database
	stmt, err := db.Prepare("INSERT INTO person (tag, wholeFame, clanStatus, joinDate, fk_clan) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}
	defer stmt.Close()

	log.Println(person)
	log.Println("person")
	result, err := stmt.Exec(person.Tag, person.WholeFame, person.ClanStatus, person.JoinDate, person.FkClan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d Person(en) erstellt", rowsAffected)})
}

//--------------------------------------------- Update ---------------------------------------------

// Update Person in Database
func updatePerson(c *gin.Context) {
	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}
	defer db.Close()

	// Parse request body
	var person Person
	err = c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültige Anfrage"})
		return
	}

	// Update person in the database
	stmt, err := db.Prepare("UPDATE person SET wholeFame = ?, clanStatus = ?, joinDate = ?, fk_clan = ? WHERE tag = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(person.WholeFame, person.ClanStatus, person.JoinDate, person.FkClan, person.Tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d Person(en) aktualisiert", rowsAffected)})
}

//--------------------------------------------- Delete ---------------------------------------------

// Delete Person from Database
func deletePerson(c *gin.Context) {
	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}
	defer db.Close()

	// Parse request body
	var person Person
	err = c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültige Anfrage"})
		return
	}

	// Delete person from the database
	stmt, err := db.Prepare("DELETE FROM person WHERE tag = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(person.Tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d Person(en) gelöscht", rowsAffected)})
}
