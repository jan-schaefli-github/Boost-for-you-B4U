package drt_person

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"database/sql"
)


// Create a new daily report
func CreateDailyReport(fameToday int, decksUsedToday int, missedDecksToday, repairPointsToday int, boatAttacksToday int, dayIdentifier string, fk_person string) {

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
	stmt, err := db.Prepare("INSERT INTO daily_report (fameToday, decksUsedToday, missedDecksToday, repairPointsToday, boatAttacksToday, dayIdentifier, fk_person) VALUES (?, ?, ?, ?, ?, ?, ?)")
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
	_, err = stmt.Exec(fameToday, decksUsedToday, missedDecksToday, repairPointsToday, boatAttacksToday, dayIdentifier, fk_person)
	if err != nil {
		logger.LogMessage("Database", "Error while executing statement: "+err.Error())
		return
	}

}

// Update a daily report
func UpdateDailyReport(fameToday int, decksUsedToday int, missedDecksToday int, repairPointsToday int, boatAttacksToday int, dayIdentifier string, fk_person string) {

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
	stmt, err := db.Prepare("UPDATE daily_report SET fameToday = ?, decksUsedToday = ?, missedDecksToday = ?, repairPointsToday = ?, boatAttacksToday = ? WHERE dayIdentifier = ? AND fk_person = ?")
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
	_, err = stmt.Exec(fameToday, decksUsedToday, missedDecksToday, repairPointsToday, boatAttacksToday, dayIdentifier, fk_person)
	if err != nil {
		logger.LogMessage("Database", "Error while executing statement: "+err.Error())
		return
	}
}


// Get the last daily report for a person
func GetLastDailyReport(fk_person string, dayIdentifier string) (int, int, int, int, int) {
	// Connect to the database
	db, err := tools.ConnectToDatabase()
	if err != nil {
		logger.LogMessage("Database", "Error while connecting to database: "+err.Error())
		return 0, 0, 0, 0, 0
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
	stmt, err := db.Prepare("SELECT fameToday, decksUsedToday, missedDecksToday, repairPointsToday, boatAttacksToday FROM daily_report WHERE fk_person = ? AND dayIdentifier != ? ORDER BY date DESC, id DESC LIMIT 1;")
	if err != nil {
		logger.LogMessage("Database", "Error while preparing statement: "+err.Error())
		return 0, 0, 0, 0, 0
	}

	defer func() {
		err := stmt.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing statement: "+err.Error())
			return
		}
	}()

	// Execute the statement and retrieve the last daily report fields
	var fameTodayNullable sql.NullInt64
	var decksUsedTodayNullable sql.NullInt64
	var missedDecksTodayNullable sql.NullInt64
	var repairPointsTodayNullable sql.NullInt64
	var boatAttacksTodayNullable sql.NullInt64
	err = stmt.QueryRow(fk_person, dayIdentifier).Scan(&fameTodayNullable, &decksUsedTodayNullable, &missedDecksTodayNullable, &repairPointsTodayNullable, &boatAttacksTodayNullable)
	if err != nil {
		if err == sql.ErrNoRows {
			// No daily_report found for fk_person
			return 0, 0, 0, 0, 0
		}
		logger.LogMessage("Database", "Error while executing statement: "+err.Error())
		return 0, 0, 0, 0, 0
	}

		fameYesterday := int(fameTodayNullable.Int64)

		decksUsedYesterday := int(decksUsedTodayNullable.Int64)

		missedDecksYesterday := int(missedDecksTodayNullable.Int64)

		repairPointsYesterday := int(repairPointsTodayNullable.Int64)

		boatAttacksYesterday := int(boatAttacksTodayNullable.Int64)

	return fameYesterday, decksUsedYesterday, missedDecksYesterday, repairPointsYesterday, boatAttacksYesterday

}


// Get the last daily report for a person
func GetLastDayIdentifier(fk_person string) (string) {
	// Connect to the database
	db, err := tools.ConnectToDatabase()
	if err != nil {
		logger.LogMessage("Database", "Error while connecting to database: "+err.Error())
		return ""
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
	stmt, err := db.Prepare("SELECT dayIdentifier FROM daily_report WHERE fk_person = ? ORDER BY date DESC, id DESC LIMIT 1;")
	if err != nil {
		logger.LogMessage("Database", "Error while preparing statement: "+err.Error())
		return ""
	}

	defer func() {
		err := stmt.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing statement: "+err.Error())
			return
		}
	}()

	// Execute the statement and retrieve the last daily report fields
	var dayIdentifierNullable sql.NullString
	err = stmt.QueryRow(fk_person).Scan(&dayIdentifierNullable)
	if err != nil {
		if err == sql.ErrNoRows {
			// No daily_report found for fk_person
			return ""
		}
		logger.LogMessage("Database", "Error while executing statement: "+err.Error())
		return ""
	}

		lastDayIdentifier := dayIdentifierNullable.String

	return lastDayIdentifier

}