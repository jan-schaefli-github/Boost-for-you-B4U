package dep_person

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPerson(c *gin.Context) {
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

	var person tools.Person
	err = db.QueryRow("SELECT * FROM person WHERE tag = '#2Y9VQVJ8'").Scan(&person.Tag, &person.WholeFame, &person.ClanStatus, &person.JoinDate, &person.FkClan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while querying database: "+err.Error())
		return
	}

	personJSON, err := json.Marshal(person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logger.LogMessage("Database", "Error while marshalling person: "+err.Error())
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(personJSON))
}
