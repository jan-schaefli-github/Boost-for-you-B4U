package dep_clan

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetClanWarlog returns war data from weekly_report, person, and daily_report tables
func GetClanWarlog(c *gin.Context) {
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

	query := `SELECT p.tag, p.name, p.clanStatus, dr.fame, wr.missedDecks, dr.decksUsedToday, dr.date
			  FROM weekly_report wr INNER JOIN person p ON wr.fk_person = p.tag INNER JOIN ( SELECT dr1.* FROM daily_report dr1 WHERE dr1.date = ( SELECT MAX(dr2.date) - ? FROM daily_report dr2 WHERE dr1.fk_person = dr2.fk_person) ) dr ON dr.fk_person = p.tag AND wr.weekIdentifier = LEFT(dr.dayIdentifier, 3)
			  WHERE p.fk_clan = ?;`

	daySubtract := 0
	fk_clan := "#P9UVQCJV"

	rows, err := db.Query(query, daySubtract, fk_clan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while querying the database: "+err.Error())
		return
	}
	defer rows.Close()

	var warData []tools.WarData

	for rows.Next() {
		var rowData tools.WarData
		err := rows.Scan(
			&rowData.Tag,
			&rowData.Name,
			&rowData.ClanStatus,
			&rowData.Fame,
			&rowData.MissedDecks,
			&rowData.DecksUsedToday,
			&rowData.Date,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
			logger.LogMessage("Database", "Error while scanning rows: "+err.Error())
			return
		}
		warData = append(warData, rowData)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while iterating over rows: "+err.Error())
		return
	}

	warDataJSON, err := json.Marshal(warData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while marshalling war data: "+err.Error())
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(warDataJSON))
}
