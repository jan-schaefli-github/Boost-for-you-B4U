package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

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

// Create or Update Person in Database
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

// Create or update daily report in the database
func createDailyReport(decksUsedToday float64, fame float64, dayIndex int, fk_person string) {
	db, err := connectToDatabase()
	if err != nil {
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

	// Check if an entry with the same dayIndex and today's date or yesterday's date already exists
	today := time.Now().Format("2006-01-02")
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM daily_report WHERE (dayIndex = ? AND date = ? AND fk_person = ?) OR (dayIndex = ? AND date = ? AND fk_person = ?)", dayIndex, today, fk_person, dayIndex, yesterday, fk_person).Scan(&count)
	if err != nil {
		logMessage("Database", "Error while querying database: "+err.Error())
		return
	}

	// Update the existing entry if it meets the conditions
	if count > 0 {
		stmt, err := db.Prepare("UPDATE daily_report SET decksUsedToday = ?, fame = ? WHERE ((dayIndex = ? AND date = ?) OR (dayIndex = ? AND date = ?)) AND fk_person = ? AND (decksUsedToday < ? OR fame < ?)")
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

		_, err = stmt.Exec(decksUsedToday, fame, dayIndex, today, dayIndex, yesterday, fk_person, decksUsedToday, fame)
		if err != nil {
			logMessage("Database", "Error while executing statement: "+err.Error())
			return
		}

		// logMessage("Database", "Daily report updated for dayIndex "+fmt.Sprint(dayIndex))
	} else {

		// Create a new entry if no matching entry exists
		stmt, err := db.Prepare("INSERT INTO daily_report (decksUsedToday, fame, dayIndex, date, fk_person) VALUES (?, ?, ?, ?, ?)")
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

		_, err = stmt.Exec(decksUsedToday, fame, dayIndex, today, fk_person)
		if err != nil {
			logMessage("Database", "Error while executing statement: "+err.Error())
			return
		}

		// logMessage("Database", "New daily report created for dayIndex "+fmt.Sprint(dayIndex))
	}
}

// Create or update weekly report in the database
func createWeeklyReport(decksUsedThisWeek float64, fame float64, weekIndex int, fk_person string) {

}

//--------------------------------------------- Check ---------------------------------------------

// Check if a person exists in the database
func checkPerson(tag string) bool {

	// Connect to the database
	db, err := connectToDatabase()
	if err != nil {
		logMessage("Database", "Error while connecting to database: "+err.Error())
		return false
	}

	// Close the database connection
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logMessage("Database", "Error while closing database: "+err.Error())
			return
		}
	}(db)

	// Check if the person exists
	var person Person
	err = db.QueryRow("SELECT * FROM person WHERE tag = ?", tag).Scan(&person.Tag, &person.WholeFame, &person.ClanStatus, &person.JoinDate, &person.FkClan)
	if err != nil {
		return false
	}

	return true
}
