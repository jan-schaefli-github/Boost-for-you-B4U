package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get Clan Information
func getClanHandler(c *gin.Context) {
	url := "https://api.clashroyale.com/v1/clans/" + encodedClanTag
	response, err := makeRequest(url)
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
	url := "https://api.clashroyale.com/v1/clans/" + encodedClanTag + "/members"
	response, err := makeRequest(url)
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
	url := "https://api.clashroyale.com/v1/clans/" + encodedClanTag + "/currentriverrace"
	response, err := makeRequest(url)
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
	url := "https://api.clashroyale.com/v1/clans/" + encodedClanTag + "/riverracelog"
	response, err := makeRequest(url)
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
func getClanMemberLeaderboardHandler(c *gin.Context) {
	url := "https://api.clashroyale.com/v1/clans/" + encodedClanTag + "/members"
	response, err := makeRequest(url)
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

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logMessage("Request", "Error while decoding response: "+err.Error())
		return
	}

	// Extract the desired values
	var result []map[string]interface{}
	for _, member := range data["items"].([]interface{}) {
		memberMap := member.(map[string]interface{})
		result = append(result, map[string]interface{}{
			"id":       memberMap["tag"],
			"name":     memberMap["name"],
			"role":     memberMap["role"],
			"clanRank": memberMap["clanRank"],
		})
	}

	c.JSON(http.StatusOK, result)
}

/*
====================================================================================================
	Part of Samuel's Code End
====================================================================================================
*/
