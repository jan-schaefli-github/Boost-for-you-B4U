package drt_person

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"database/sql"
)

// Create a new weekly report
func CreateWeeklyReport(weekIdentifier string, fk_person string) {

	// Connect to the database
	db, err := tools.ConnectToDatabase()
	if err != nil {
		logger.LogMessage("Database", "Error while connecting to database: "+err.Error())
		return
	}
	
	// Close the database connection
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing database: "+err.Error())
			return
		}
	}(db)

	// Insert person into the database
	stmt, err := db.Prepare("INSERT INTO weekly_report (weekIdentifier, fk_person) VALUES (?, ?)")
	if err != nil {
		logger.LogMessage("Database", "Error while preparing statement: "+err.Error())
		return
	}

	// Close the statement
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing statement: "+err.Error())
			return
		}
	}(stmt)

	// Execute the statement
	_, err = stmt.Exec(weekIdentifier, fk_person)
	if err != nil {
		logger.LogMessage("Database", "Error while executing statement: "+err.Error())
		return
	}

}

// Update a weekly report
func UpdateWeeklyReport(fame int, lastMissedDecks int, weekIdentifier string, fk_person string) {

	// Connect to the database
	db, err := tools.ConnectToDatabase()
	if err != nil {
		logger.LogMessage("Database", "Error while connecting to database: "+err.Error())
		return
	}
	
	// Close the database connection
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing database: "+err.Error())
			return
		}
	}(db)

	// Insert person into the database
	stmt, err := db.Prepare("UPDATE weekly_report SET fame = ?, missedDecks = ? WHERE weekIdentifier = ? AND fk_person = ?")
	if err != nil {
		logger.LogMessage("Database", "Error while preparing statement: "+err.Error())
		return
	}

	// Close the statement
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing statement: "+err.Error())
			return
		}
	}(stmt)

	// Execute the statement
	_, err = stmt.Exec(fame, lastMissedDecks, weekIdentifier, fk_person)
	if err != nil {
		logger.LogMessage("Database", "Error while executing statement: "+err.Error())
		return
	}

}

// Get the last weekly report for a person
func GetLastWeeklyReport(fk_person string) (lastMissedDecks int, lastWeekIdentifier string) {
	// Connect to the database
	db, err := tools.ConnectToDatabase()
	if err != nil {
		logger.LogMessage("Database", "Error while connecting to database: "+err.Error())
		return 0, ""
	}

	// Close the database connection
	defer func() {
		err := db.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing database: "+err.Error())
			return
		}
	}()

	// Prepare the statement
	stmt, err := db.Prepare("SELECT missedDecks, weekIdentifier FROM weekly_report WHERE fk_person = ? ORDER BY date DESC LIMIT 1")
	if err != nil {
		logger.LogMessage("Database", "Error while preparing statement: "+err.Error())
		return 0, ""
	}

	defer func() {
		err := stmt.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing statement: "+err.Error())
			return
		}
	}()

	// Execute the statement and retrieve the last weekly report fields
	var lastMissedDecksNullable sql.NullInt64
	var lastWeekIdentifierNullable sql.NullString
	err = stmt.QueryRow(fk_person).Scan(&lastMissedDecksNullable, &lastWeekIdentifierNullable)
	if err != nil {
		if err == sql.ErrNoRows {
			// No weekly_report found for fk_person
			return 0, ""
		}
		logger.LogMessage("Database", "Error while executing statement: "+err.Error())
		return 0, ""
	}

	if lastMissedDecksNullable.Valid {
		lastMissedDecks = int(lastMissedDecksNullable.Int64)
	}

	if lastWeekIdentifierNullable.Valid {
		lastWeekIdentifier = lastWeekIdentifierNullable.String
	}

	return lastMissedDecks, lastWeekIdentifier
}
