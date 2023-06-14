package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
)

// Get Clan Information
func getClanHandler(c *gin.Context) {
	urlForApiRequest := "https://api.clashroyale.com/v1/clans/" + encodedClanTag
	response, err := makeRequest(urlForApiRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logMessage("Request", "Error while making request: "+err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logMessage("Request", "Error while closing response body: "+err.Error())
			return
		}
	}(response.Body)

	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logMessage("Request", "Error while decoding response: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

// Get Members of Clan
func getMembersHandler(c *gin.Context) {
	urlForApiRequest := "https://api.clashroyale.com/v1/clans/" + encodedClanTag + "/members"
	response, err := makeRequest(urlForApiRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logMessage("Request", "Error while making request: "+err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logMessage("Request", "Error while decoding response: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

// Get Current Ri verrace information
func getCurrentRiverRaceHandler(c *gin.Context) {
	urlForApiRequest := "https://api.clashroyale.com/v1/clans/" + encodedClanTag + "/currentriverrace"
	response, err := makeRequest(urlForApiRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logMessage("Request", "Error while making request: "+err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logMessage("Request", "Error while closing response body: "+err.Error())
		}
	}(response.Body)

	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logMessage("Request", "Error while decoding response: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

func getRiverRaceLogHandler(c *gin.Context) {
	urlForApiRequest := "https://api.clashroyale.com/v1/clans/" + encodedClanTag + "/riverracelog"
	response, err := makeRequest(urlForApiRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logMessage("Request", "Error while making request: "+err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logMessage("Request", "Error while closing response body: "+err.Error())
		}
	}(response.Body)

	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logMessage("Request", "Error while decoding response: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

func makeRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Access-Control-Allow-Origin", "*")

	client := &http.Client{}
	return client.Do(req)
}

/*
====================================================================================================

	Part of Samuel's Code Start

====================================================================================================
*/

/* Endpoint for Clan Member leaderboard encodedClanTag := url.QueryEscape(clanTag)*/

func getClanMemberLeaderboardHandler(c *gin.Context) {
	// Abfragen der Parameter
	clanID := url.QueryEscape(c.Query("clanID"))

	// Erstellen der API-Anforderungs-URL
	urlForApiRequest := "https://api.clashroyale.com/v1/clans/" + clanID + "/members"

	// API-Anforderung durchführen
	response, err := makeRequest(urlForApiRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist beim Senden der API-Anfrage aufgetreten"})
		logMessage("Request", "Fehler beim Senden der Anfrage: "+err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		// Antwort-Body schließen und mögliche Fehler behandeln
		err := Body.Close()
		if err != nil {
			logMessage("Request", "Fehler beim Schließen des Antwort-Bodys: "+err.Error())
		}
	}(response.Body)

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist beim Dekodieren der API-Antwort aufgetreten"})
		logMessage("Request", "Fehler beim Dekodieren der Antwort: "+err.Error())
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

/* Endpoint for Locations */

func getLocationsHandler(c *gin.Context) {

	// Festlegen der API-Anforderungs-URL
	urlForApiRequest := "https://api.clashroyale.com/v1/locations"

	// API-Anforderung durchführen
	response, err := makeRequest(urlForApiRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist beim Senden der API-Anfrage aufgetreten"})
		logMessage("Request", "Fehler beim Senden der Anfrage: "+err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		// Antwort-Body schließen und mögliche Fehler behandeln
		err := Body.Close()
		if err != nil {
			logMessage("Request", "Fehler beim Schließen des Antwort-Bodys: "+err.Error())
		}
	}(response.Body)

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist beim Dekodieren der API-Antwort aufgetreten"})
		logMessage("Request", "Fehler beim Dekodieren der Antwort: "+err.Error())
		return
	}
	println(data)
	c.JSON(http.StatusOK, data)
}

/* Endpoint for Clans by Location */
func getRankingByLocationHandler(c *gin.Context) {
	// Abfragen der Parameter
	locationID := url.QueryEscape(c.Query("locationID"))

	// Erstellen der API-Anforderungs-URL + Setzen des Limits auf 100
	urlForApiRequest := "https://api.clashroyale.com/v1/locations/" + locationID + "/rankings/clans?limit=100"

	// API-Anforderung durchführen
	response, err := makeRequest(urlForApiRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist beim Senden der API-Anfrage aufgetreten"})
		logMessage("Request", "Fehler beim Senden der Anfrage: "+err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		// Antwort-Body schließen und mögliche Fehler behandeln
		err := Body.Close()
		if err != nil {
			logMessage("Request", "Fehler beim Schließen des Antwort-Bodys: "+err.Error())
		}
	}(response.Body)

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist beim Dekodieren der API-Antwort aufgetreten"})
		logMessage("Request", "Fehler beim Dekodieren der Antwort: "+err.Error())
		return
	}
	println(data)
	c.JSON(http.StatusOK, data)
}

/*
====================================================================================================
	Part of Samuel's Code End
====================================================================================================
*/
