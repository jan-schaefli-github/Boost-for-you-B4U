package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

//--------------------------------------------- Get ---------------------------------------------

// Get Clan from Database
func getClan(c *gin.Context) {
	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logMessage("Database", "Error while connecting to database: "+err.Error())
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logMessage("Database", "Error while closing database: "+err.Error())
			return
		}
	}(db)

	var clan Clan
	err = db.QueryRow("SELECT * FROM clan WHERE tag = '#P9UVQCJV'").Scan(&clan.Tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logMessage("Database", "Error while querying database: "+err.Error())
		return
	}

	clanJSON, err := json.Marshal(clan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logMessage("Database", "Error while marshalling clan: "+err.Error())
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
		logMessage("Database", "Error while connecting to database: "+err.Error())
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logMessage("Database", "Error while closing database: "+err.Error())
			return
		}
	}(db)

	var person Person
	err = db.QueryRow("SELECT * FROM person WHERE tag = '#2Y9VQVJ8'").Scan(&person.Tag, &person.WholeFame, &person.ClanStatus, &person.JoinDate, &person.FkClan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logMessage("Database", "Error while querying database: "+err.Error())
		return
	}

	personJSON, err := json.Marshal(person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logMessage("Database", "Error while marshalling person: "+err.Error())
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(personJSON))
}

// Get Clan Tags from Database
func getClanTags() []string {
	db, err := connectToDatabase()
	if err != nil {
		logMessage("Database", "Error while connecting to database: "+err.Error())
		return nil
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logMessage("Database", "Error while closing database: "+err.Error())
			return
		}
	}(db)

	rows, err := db.Query("SELECT tag FROM clan")
	if err != nil {
		logMessage("Database", "Error while querying database: "+err.Error())
		return nil
	}
	defer rows.Close()

	var tags []string

	for rows.Next() {
		var tag string
		err := rows.Scan(&tag)
		if err != nil {
			logMessage("Database", "Error while scanning rows: "+err.Error())
			return nil
		}
		tags = append(tags, tag)
	}

	err = rows.Err()
	if err != nil {
		logMessage("Database", "Error while scanning rows: "+err.Error())
		return nil
	}

	return tags
}

//--------------------------------------------- Create ---------------------------------------------

// Create Person in Database
func createPerson(tag string, fk_clan string) {
	db, err := connectToDatabase()
	if err != nil {
		logMessage("Database", "Error while connecting to database: "+err.Error())
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logMessage("Database", "Error while closing database: "+err.Error())
		}
	}(db)

	// Insert person into the database
	stmt, err := db.Prepare("INSERT INTO person (tag, fk_clan) VALUES (?, ?)")
	if err != nil {
		logMessage("Database", "Error while preparing statement: "+err.Error())
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			logMessage("Database", "Error while closing statement: "+err.Error())
			return
		}
	}(stmt)

	// Execute the statement
	_, err = stmt.Exec(tag, fk_clan)
	if err != nil {
		logMessage("Database", "Error while executing statement: "+err.Error())
		return
	}

	logMessage("Database", "Person with tag "+tag+" created")
}

//--------------------------------------------- Check ---------------------------------------------

func checkPerson(tag string) bool {
	db, err := connectToDatabase()
	if err != nil {
		logMessage("Database", "Error while connecting to database: "+err.Error())
		return false
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logMessage("Database", "Error while closing database: "+err.Error())
			return
		}
	}(db)

	var person Person
	err = db.QueryRow("SELECT * FROM person WHERE tag = ?", tag).Scan(&person.Tag, &person.WholeFame, &person.ClanStatus, &person.JoinDate, &person.FkClan)
	if err != nil {
		return false
	}

	return true
}
