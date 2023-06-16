package drt_person

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"database/sql"
)


// Create a new daily report
func CreateDailyReport(decksUsedToday float64, fame float64, dayIdentifier string, fk_person string) {

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
	stmt, err := db.Prepare("INSERT INTO daily_report (decksUsedToday, fame, dayIdentifier, fk_person) VALUES (?, ?, ?, ?)")
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
	_, err = stmt.Exec(decksUsedToday, fame, dayIdentifier, fk_person)
	if err != nil {
		logger.LogMessage("Database", "Error while executing statement: "+err.Error())
		return
	}

}

// Update a daily report
func UpdateDailyReport(decksUsedToday float64, fame float64, dayIdentifier string, fk_person string) {

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
	stmt, err := db.Prepare("UPDATE daily_report SET decksUsedToday = ?, fame = ? WHERE dayIdentifier = ? AND fk_person = ?")
	if err != nil {
		logger.LogMessage("Database", "Error while preparing statement: "+err.Error())
		return
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing statement: "+err.Error())
			return
		}
	}(stmt)

	// Execute the statement
	_, err = stmt.Exec(decksUsedToday, fame, dayIdentifier, fk_person)
	if err != nil {
		logger.LogMessage("Database", "Error while executing statement: "+err.Error())
		return
	}
}


// Get the last daily report for a person
func GetLastDailyReport(fk_person string) (decksUsedYesterday int, lastFame int, lastDayIdentifier string) {
	// Connect to the database
	db, err := tools.ConnectToDatabase()
	if err != nil {
		logger.LogMessage("Database", "Error while connecting to database: "+err.Error())
		return 0, 0, ""
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
	stmt, err := db.Prepare("SELECT decksUsedToday, fame, dayIdentifier FROM daily_report WHERE fk_person = ? ORDER BY date DESC, id DESC LIMIT 1;")
	if err != nil {
		logger.LogMessage("Database", "Error while preparing statement: "+err.Error())
		return 0, 0, ""
	}

	defer func() {
		err := stmt.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing statement: "+err.Error())
			return
		}
	}()

	// Execute the statement and retrieve the last daily report fields
	var decksUsedYesterdayNullable sql.NullInt64
	var lastFameNullable sql.NullInt64
	var lastDayIdentifierNullable sql.NullString
	err = stmt.QueryRow(fk_person).Scan(&decksUsedYesterdayNullable, &lastFameNullable, &lastDayIdentifierNullable)
	if err != nil {
		if err == sql.ErrNoRows {
			// No daily_report found for fk_person
			return 0, 0, ""
		}
		logger.LogMessage("Database", "Error while executing statement: "+err.Error())
		return 0, 0, ""
	}

	if decksUsedYesterdayNullable.Valid {
		decksUsedYesterday = int(decksUsedYesterdayNullable.Int64)
	}

	if lastFameNullable.Valid {
		lastFame = int(lastFameNullable.Int64)
	}

	if lastDayIdentifierNullable.Valid {
		lastDayIdentifier = lastDayIdentifierNullable.String
	}

	return decksUsedYesterday, lastFame, lastDayIdentifier
}

