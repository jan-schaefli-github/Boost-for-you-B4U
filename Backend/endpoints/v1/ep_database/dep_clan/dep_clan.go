package dep_clan

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Get Clan from Database
func GetClan(c *gin.Context) {
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

	var clan tools.Clan
	err = db.QueryRow("SELECT * FROM clan WHERE tag = '#P9UVQCJV'").Scan(&clan.Tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while querying database: "+err.Error())
		return
	}

	clanJSON, err := json.Marshal(clan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while marshalling clan: "+err.Error())
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(clanJSON))
}
