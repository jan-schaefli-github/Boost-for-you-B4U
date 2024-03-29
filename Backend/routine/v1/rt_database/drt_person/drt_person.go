package drt_person

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"database/sql"
	"strings"
)

func CreatePerson(tag string, name string, role string, trophies int, clanRank int, fk_clan string) {
	
	// Connect to the database
	db, err := tools.ConnectToDatabase()
	if err != nil {
		logger.LogMessage("Database", "Error while connecting to database: "+err.Error())
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing database: "+err.Error())
			return
		}
	}(db)

	// Create person in the database
	stmt, err := db.Prepare("INSERT INTO person(tag, name, role, trophies, clanRank, fk_clan) VALUES(?, ?, ?, ?, ?, ?)")
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
	_, err = stmt.Exec(tag, name, role, trophies, clanRank, fk_clan)
	if err != nil {
		logger.LogMessage("Database", "Error while executing statement: "+err.Error())
		return
	}

	logger.LogMessage("Database", "Person joined: "+tag)
}

func UpdatePerson(tag string, name string, role string, trophies int, clanRank int, wholeFame int, wholeDecksUsed int, wholeMissedDecks int, wholeRepairPoints int, wholeBoatAttacks int) {
	
	// Connect to the database
	db, err := tools.ConnectToDatabase()
	if err != nil {
		logger.LogMessage("Database", "Error while connecting to database: "+err.Error())
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing database: "+err.Error())
			return
		}
	}(db)

	// Update person in the database
	stmt, err := db.Prepare("UPDATE person SET name = ?, role = ?, trophies = ?, clanRank = ?, wholeFame = ?, wholeDecksUsed = ?, wholeMissedDecks = ?, wholeRepairPoints = ?, wholeBoatAttacks = ? WHERE tag = ?")
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
	_, err = stmt.Exec(name, role, trophies, clanRank, wholeFame, wholeDecksUsed, wholeMissedDecks, wholeRepairPoints, wholeBoatAttacks, tag)
	if err != nil {
		logger.LogMessage("Database", "Error while executing statement: "+err.Error())
		return
	}

	//logger.LogMessage("Database", "Person updated: "+tag)
}

func GetPerson(tag string) (int, int, int, int, int) {

	// Connect to the database
	db, err := tools.ConnectToDatabase()
	if err != nil {
		logger.LogMessage("Database", "Error while connecting to database: "+err.Error())
		return 0, 0, 0, 0, 0
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing database: "+err.Error())
			return
		}
	}(db)

	// Get person from the database
	var wholeFame int
	var wholeDecksUsed int
	var wholeMissedDecks int
	var wholeRepairPoints int
	var wholeBoatAttacks int
	err = db.QueryRow("SELECT wholeFame, wholeDecksUsed, wholeMissedDecks, wholeRepairPoints, wholeBoatAttacks FROM person WHERE tag = ?", tag).Scan(&wholeFame, &wholeDecksUsed, &wholeMissedDecks, &wholeRepairPoints, &wholeBoatAttacks)
	if err != nil {
		logger.LogMessage("Database", "Error while getting person: "+err.Error())
		return 0, 0, 0, 0, 0
	}

	return wholeFame, wholeDecksUsed, wholeMissedDecks, wholeRepairPoints, wholeBoatAttacks
}

	
// Check if a person exists in the database
func CheckPerson(tag string) bool {
	// Connect to the database
	db, err := tools.ConnectToDatabase()
	if err != nil {
		logger.LogMessage("Database", "Error while connecting to database: "+err.Error())
		return false
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing database: "+err.Error())
			return
		}
	}(db)

	// Check if the person exists
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM person WHERE tag = ?", tag).Scan(&count)
	if err != nil {
		logger.LogMessage("Database", "Error while checking if person exists: "+err.Error())
		return false
	}

	if count > 0 {
		return true // Person exists
	}

	return false // Person does not exist
}

// Routine to update the status of the persons
func UpdatePersonStatus(members []string, clanTag string) {

	// Connect to database
	db, err := tools.ConnectToDatabase()
	if err != nil {
		logger.LogMessage("Routine", "Error while connecting to database: "+err.Error())
		return
	}

	// Close database connection
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.LogMessage("Routine", "Error while closing database: "+err.Error())
			return
		}
	}(db)

	// Create a string of placeholders for the query
	placeholders := strings.Repeat("?, ", len(members)-1) + "?"

	// Create the query
	query := "UPDATE person SET clanStatus = CASE " +
		"WHEN tag IN (" + placeholders + ") THEN 1 " +
		"ELSE 0 END" +
		" WHERE fk_clan = '" + clanTag + "'"

	// Prepare the query
	stmt, err := db.Prepare(query)
	if err != nil {
		logger.LogMessage("Routine", "Error while preparing query: "+err.Error())
		return
	}

	// Close the statement
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			logger.LogMessage("Routine", "Error while closing statement: "+err.Error())
			return
		}
	}(stmt)

	// Create a slice of interface{} with the length of members
	args := make([]interface{}, len(members))
	for i, player := range members {
		args[i] = player
	}

	// Execute the query
	_, err = stmt.Exec(args...)
	if err != nil {
		logger.LogMessage("Routine", "Error while executing query: "+err.Error())
		return
	}

	return
}

