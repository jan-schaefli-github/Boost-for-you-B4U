package dep_person

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetDailyReport returns all daily reports
func GetDailyReport(c *gin.Context) {
	db, err := tools.ConnectToDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while connecting to database: " + err.Error())
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing database: " + err.Error())
			return
		}
	}(db)

	rows, err := db.Query("SELECT * FROM daily_report")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while querying database: " + err.Error())
		return
	}
	defer rows.Close()

	var dailyReports []tools.DailyReport
	for rows.Next() {
		var dailyReport tools.DailyReport
		err := rows.Scan(&dailyReport.ID, &dailyReport.DecksUsedToday, &dailyReport.Fame, &dailyReport.DayIdentifier, &dailyReport.Date, &dailyReport.FkPerson)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
			logger.LogMessage("Database", "Error while scanning rows: " + err.Error())
			return
			}
		dailyReports = append(dailyReports, dailyReport)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while iterating over rows: " + err.Error())
		return
	}

	dailyReportsJSON, err := json.Marshal(dailyReports)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while marshalling daily reports: " + err.Error())
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(dailyReportsJSON))
}
