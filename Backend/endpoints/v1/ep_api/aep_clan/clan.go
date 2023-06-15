package aep_clan

import (
	"b4u/backend/logs"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func getClanHandler(c *gin.Context) {
	urlForApiRequest := "https://api.clashroyale.com/v1/clans/" + encodedClanTag
	response, err := makeRequest(urlForApiRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logs.logMessage("Request", "Error while making request: "+err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logs.logMessage("Request", "Error while closing response body: "+err.Error())
			return
		}
	}(response.Body)

	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		logs.logMessage("Request", "Error while decoding response: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}
