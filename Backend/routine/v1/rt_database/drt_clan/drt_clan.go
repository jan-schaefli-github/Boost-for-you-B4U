package drt_clan

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"database/sql"
)

// GetClanTags Get Clan Tags from Database
func GetClanTags() []string {
	db, err := tools.ConnectToDatabase()
	if err != nil {
		logger.LogMessage("Database", "Error while connecting to database: "+err.Error())
		return nil
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing database: "+err.Error())
			return
		}
	}(db)

	rows, err := db.Query("SELECT tag FROM clan")
	if err != nil {
		logger.LogMessage("Database", "Error while querying database: "+err.Error())
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing rows: "+err.Error())
			return
		}
	}(rows)

	var tags []string

	for rows.Next() {
		var tag string
		err := rows.Scan(&tag)
		if err != nil {
			logger.LogMessage("Database", "Error while scanning rows: "+err.Error())
			return nil
		}
		tags = append(tags, tag)
	}

	err = rows.Err()
	if err != nil {
		logger.LogMessage("Database", "Error while scanning rows: "+err.Error())
		return nil
	}

	return tags
}
