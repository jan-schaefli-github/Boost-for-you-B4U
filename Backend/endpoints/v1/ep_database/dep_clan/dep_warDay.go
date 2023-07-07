package dep_clan

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get day log returns war data from person and daily_report tables
func GetClanDayLog(c *gin.Context) {
	db, err := tools.ConnectToDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while connecting to the database: "+err.Error())
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing the database: "+err.Error())
			return
		}
	}(db)

	// Execute the "SET sql_mode = ''" statement
	_, err = db.Exec("SET sql_mode = '';")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while executing 'SET sql_mode = ''': "+err.Error())
		return
	}

	query := `SELECT p.tag, p.name, p.joinDate, p.clanStatus, p.role, p.trophies, p.clanRank, dr.fameToday, dr.decksUsedToday, dr.missedDecksToday, dr.repairPointsToday, dr.boatAttacksToday
			FROM person p
			INNER JOIN daily_report dr ON p.tag = dr.fk_person
			WHERE p.fk_clan = ? AND dr.dayIdentifier = (SELECT dr2.dayIdentifier FROM daily_report dr2 INNER JOIN person p2 ON p2.tag = dr2.fk_person WHERE p2.fk_clan = ? GROUP BY dr2.dayIdentifier ORDER BY dr2.dayIdentifier DESC LIMIT ?, 1);`

	tools.LoadDotEnv()
	
	fk_clan := c.Param("clanID")
	if fk_clan == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Keine Clan-ID angegeben"})
		logger.LogMessage("Database", "No clan ID specified")
		return
	} else if fk_clan == "standard" {
		fk_clan = tools.ClanTag
	} else {
		fk_clan = "#"+fk_clan
	}
	
	// Get the current date
	offset := c.Param("offset")

	if offset < "0" {
		offset = "0"
	}

	rows, err := db.Query(query, fk_clan, fk_clan, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while querying the database: "+err.Error())
		return
	}
	defer rows.Close()

	var WarLog []tools.WarLog

	for rows.Next() {
		var rowData tools.WarLog
		err := rows.Scan(
			&rowData.Tag,
			&rowData.Name,
			&rowData.JoinDate,
			&rowData.ClanStatus,
			&rowData.Role,
			&rowData.Trophies,
			&rowData.ClanRank,
			&rowData.Fame,
			&rowData.DecksUsed,
			&rowData.MissedDecks,
			&rowData.RepairPoints,
			&rowData.BoatAttacks,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
			logger.LogMessage("Database", "Error while scanning rows: "+err.Error())
			return
		}
		WarLog = append(WarLog, rowData)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while iterating over rows: "+err.Error())
		return
	}

	if len(WarLog) == 0 {
		c.JSON(http.StatusOK, gin.H{"error": "notFound"})
		return
	}

	WarLogJSON, err := json.Marshal(WarLog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while marshalling war data: "+err.Error())
		return
	}


	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(WarLogJSON))
}
