package dep_clan

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get day log returns war data from person and weekly_report tables
func GetClanWeekLog(c *gin.Context) {
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

	query := `SELECT p.tag, p.name, p.joinDate, p.clanStatus, p.role, p.trophies, p.clanRank, wr.fameThisWeek, wr.decksUsedThisWeek, wr.missedDecksThisWeek, wr.repairPointsThisWeek, wr.boatAttacksThisWeek
			FROM person p
			INNER JOIN weekly_report wr ON p.tag = wr.fk_person
			WHERE p.fk_clan = ? AND wr.weekIdentifier = (SELECT wr2.weekIdentifier FROM weekly_report wr2 INNER JOIN person p2 ON p2.tag = wr2.fk_person WHERE p2.fk_clan = ? GROUP BY wr2.weekIdentifier ORDER BY wr2.weekIdentifier DESC LIMIT ?, 1);`

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

	maxOffset := GetWeeklyReportAmount(fk_clan)
	
	// Get the current date
	offset := c.Param("offset")

	if offset < "0" {
		offset = "0"
	} else if offset > maxOffset {
		offset = maxOffset
	}

	rows, err := db.Query(query, fk_clan, fk_clan, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while querying the database: "+err.Error())
		return
	}
	defer rows.Close()

	var WarLogItems []tools.WarLogItems

	for rows.Next() {
		var rowData tools.WarLogItems
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
		WarLogItems = append(WarLogItems, rowData)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while iterating over rows: "+err.Error())
		return
	}

	if len(WarLogItems) == 0 {
		c.JSON(http.StatusOK, gin.H{"error": "notFound"})
		return
	}

	WarLog := tools.WarLog{
		Items: WarLogItems,
		MaxOffset: maxOffset,
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


// Get the amount of weekly reports
func GetWeeklyReportAmount(fk_person string) (string) {
	// Connect to the database
	db, err := tools.ConnectToDatabase()
	if err != nil {
		logger.LogMessage("Database", "Error while connecting to database: "+err.Error())
		return "0"
	}

	// Close the database connection
	defer func() {
		err := db.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing database: "+err.Error())
			return
		}
	}()

	// Prepare the statement
	stmt, err := db.Prepare("SELECT COUNT(wr.weekIdentifier) AS amount FROM weekly_report wr INNER JOIN person p on p.tag = wr.fk_person WHERE p.fk_clan = ? GROUP BY wr.fk_person Limit 1;")
	if err != nil {
		logger.LogMessage("Database", "Error while preparing statement: "+err.Error())
		return "0"
	}

	defer func() {
		err := stmt.Close()
		if err != nil {
			logger.LogMessage("Database", "Error while closing statement: "+err.Error())
			return
		}
	}()

	// Execute the statement and retrieve the last daily report fields
	var amountNullable sql.NullInt64
	err = stmt.QueryRow(fk_person).Scan(&amountNullable)
	if err != nil {
		if err == sql.ErrNoRows {
			// No daily_report found for fk_person
			return "0"
		}
		logger.LogMessage("Database", "Error while executing statement: "+err.Error())
		return "0"
	}

		amount := amountNullable.Int64
		amountStr := strconv.Itoa(int(amount) - 1)

	return amountStr

}