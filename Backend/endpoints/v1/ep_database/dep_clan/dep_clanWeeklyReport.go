package dep_clan

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetClanWeeklyReport(c *gin.Context) {
	db, err := tools.ConnectToDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
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

	rows, err := db.Query("SELECT id, fame, fame_gain FROM clan_weekly_report WHERE fk_clan = '#P9UVQCJV';")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while querying database: "+err.Error())
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing rows: "+err.Error())
			return
		}
	}(rows)

	var clans tools.ClanWeekReports

	for rows.Next() {
		var clan tools.ClanWeekReport
		err := rows.Scan(&clan.Id, &clan.Fame, &clan.FameGain)
		if err != nil {
			logger.LogMessage("Database", "Error while scanning row: "+err.Error())
			continue
		}
		clans = append(clans, clan)
	}
	if rows.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while iterating rows: "+rows.Err().Error())
		return
	}

	clansJSON, err := json.Marshal(clans)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while marshalling clans: "+err.Error())
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(clansJSON))
}
