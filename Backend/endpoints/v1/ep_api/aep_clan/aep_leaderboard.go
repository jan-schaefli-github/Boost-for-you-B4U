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

func GetClanMemberLeaderboard(c *gin.Context) {
	// Abfragen der Parameter
	clanID := url.QueryEscape(c.Query("clanID"))

	// Erstellen der API-Anforderungs-URL
	urlForApiRequest := "https://api.clashroyale.com/v1/clans/" + clanID + "/members"

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

	// Gewünschte Werte extrahieren
	var result []map[string]interface{}
	if items, ok := data["items"].([]interface{}); ok {
		for _, member := range items {
			if memberMap, ok := member.(map[string]interface{}); ok {
				result = append(result, map[string]interface{}{
					"id":       memberMap["tag"],
					"name":     memberMap["name"],
					"role":     memberMap["role"],
					"clanRank": memberMap["clanRank"],
				})
			}
		}
	}

	if len(result) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Keine Clan-Mitglieder gefunden"})
	} else {
		c.JSON(http.StatusOK, result)
	}
}
