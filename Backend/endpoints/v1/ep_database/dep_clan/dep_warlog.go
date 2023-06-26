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

	query := `SELECT p.tag, p.name, p.clanStatus, wr.fame, wr.missedDecks, dr.decksUsedToday, repairPoints, boatAttacks, p.joinDate
	FROM person p
	INNER JOIN weekly_report wr ON p.tag = wr.fk_person
	INNER JOIN daily_report dr ON p.tag = dr.fk_person
	INNER JOIN (SELECT fk_person, MAX(id) AS max_id FROM weekly_report GROUP BY fk_person) wr_max ON wr.fk_person = wr_max.fk_person AND wr.id = wr_max.max_id
	INNER JOIN (SELECT fk_person, MAX(id) AS max_id FROM daily_report GROUP BY fk_person) dr_max ON dr.fk_person = dr_max.fk_person AND dr.id = dr_max.max_id
	WHERE p.fk_clan = ? AND dr.date <= ?;`

	fk_clan := "#P9UVQCJV"
	date:= "2023-06-26"

	rows, err := db.Query(query, fk_clan, date)
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
			&rowData.RepairPoints,
			&rowData.BoatAttacks,
			&rowData.JoinDate,
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
