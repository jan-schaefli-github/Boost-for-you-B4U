package dep_clan

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetClans returns all clans from the database as JSON response
func GetClans(c *gin.Context) {
	db, err := tools.ConnectToDatabase()
	if err != nil {
		logger.LogMessage("Database", "Error while connecting to the database: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while connecting to the database"})
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing the database: "+err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while closing the database"})
			return
		}
	}(db)

	rows, err := db.Query("SELECT tag FROM clan")
	if err != nil {
		logger.LogMessage("Database", "Error while querying the database: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while querying the database"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing the rows: "+err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while closing the rows"})
			return
		}
	}(rows)

	var clans []tools.Clan

	for rows.Next() {
		var clan tools.Clan
		err := rows.Scan(&clan.Tag)
		if err != nil {
			logger.LogMessage("Database", "Error while scanning the rows: "+err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while scanning the rows"})
			return
		}
		clans = append(clans, clan)
	}

	err = rows.Err()
	if err != nil {
		logger.LogMessage("Database", "Error while scanning the rows: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while scanning the rows"})
		return
	}

	c.JSON(http.StatusOK, clans)
}
