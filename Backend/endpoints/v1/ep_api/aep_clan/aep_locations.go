package aep_clan

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
)

func GetLocations(c *gin.Context) {

	// Festlegen der API-Anforderungs-URL
	urlForApiRequest := "https://api.clashroyale.com/v1/locations"

	// API-Anforderung durchführen
	response, err := tools.MakeRequest(urlForApiRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist beim Senden der API-Anfrage aufgetreten"})
		logger.LogMessage("Request", "Fehler beim Senden der Anfrage: "+err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		// Antwort-Body schließen und mögliche Fehler behandeln
		err := Body.Close()
		if err != nil {
			logger.LogMessage("Request", "Fehler beim Schließen des Antwort-Bodys: "+err.Error())
		}
	}(response.Body)

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist beim Dekodieren der API-Antwort aufgetreten"})
		logger.LogMessage("Request", "Fehler beim Dekodieren der Antwort: "+err.Error())
		return
	}
	println(data)
	c.JSON(http.StatusOK, data)
}

func GetClanRankingByLocation(c *gin.Context) {
	// Abfragen der Parameter
	locationID := url.QueryEscape(c.Query("locationID"))

	// Erstellen der API-Anforderungs-URL + Setzen des Limits auf 100
	urlForApiRequest := "https://api.clashroyale.com/v1/locations/" + locationID + "/rankings/clans?limit=50"

	// API-Anforderung durchführen
	response, err := tools.MakeRequest(urlForApiRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist beim Senden der API-Anfrage aufgetreten"})
		logger.LogMessage("Request", "Fehler beim Senden der Anfrage: "+err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		// Antwort-Body schließen und mögliche Fehler behandeln
		err := Body.Close()
		if err != nil {
			logger.LogMessage("Request", "Fehler beim Schließen des Antwort-Bodys: "+err.Error())
		}
	}(response.Body)

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist beim Dekodieren der API-Antwort aufgetreten"})
		logger.LogMessage("Request", "Fehler beim Dekodieren der Antwort: "+err.Error())
		return
	}
	println(data)
	c.JSON(http.StatusOK, data)
}
