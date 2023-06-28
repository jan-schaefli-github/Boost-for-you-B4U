package dep_clan

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

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

	query := `SELECT p.tag, p.name, p.clanStatus, p.role, p.trophies, clanRank, wr.fame, wr.missedDecks, dr.decksUsedToday, repairPoints, boatAttacks, p.joinDate
	FROM person p
	INNER JOIN weekly_report wr ON p.tag = wr.fk_person
	INNER JOIN daily_report dr ON p.tag = dr.fk_person
	INNER JOIN (SELECT fk_person, MAX(id) AS max_id FROM weekly_report GROUP BY fk_person) wr_max ON wr.fk_person = wr_max.fk_person AND wr.id = wr_max.max_id
	INNER JOIN (SELECT fk_person, MAX(id) AS max_id FROM daily_report GROUP BY fk_person) dr_max ON dr.fk_person = dr_max.fk_person AND dr.id = dr_max.max_id
	WHERE p.fk_clan = ? AND dr.date >= ?;`

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
	now := time.Now()

	// Subtract one day from the current date
	oneDayAgo := now.AddDate(0, 0, -1)

	// Format the date as "2006-01-02"
	date := oneDayAgo.Format("2006-01-02")

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
			&rowData.Role,
			&rowData.Trophies,
			&rowData.ClanRank,
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
