package aep_clan

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func GetRiverRaceLog(c *gin.Context) {
	urlForApiRequest := "https://api.clashroyale.com/v1/clans/" + tools.EncodedClanTag + "/riverracelog"
	response, err := tools.MakeRequest(urlForApiRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while processing the request"})
		logger.LogMessage("Request", "Error while making request: "+err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.LogMessage("Request", "Error while closing response body: "+err.Error())
		}
	}(response.Body)

	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while decoding the response"})
		logger.LogMessage("Request", "Error while decoding response: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetRiverRaceLogChartFormatted(c *gin.Context) {
	/* Get URL Parameters */
	clanTags := c.Query("clanTag")
	locationID := url.QueryEscape(c.Query("locationID"))

	/* Process URL Parameters */
	// Split clanTags into an array.
	clanTagsArray := strings.Split(clanTags, ",")

	// Get and Split locationId Tag of clans into an array.
	// Create API request URL + Set the limit to 10
	locationIDURLForAPIRequest := "https://api.clashroyale.com/v1/locations/" + locationID + "/rankings/clans?limit=10"

	// Perform API request
	response, err := tools.MakeRequest(locationIDURLForAPIRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while sending the API request"})
		logger.LogMessage("Request", "Error sending the request: "+err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		// Close the response body and handle possible errors
		err := Body.Close()
		if err != nil {
			logger.LogMessage("Request", "Error closing the response body: "+err.Error())
		}
	}(response.Body)

	var data map[string]interface{} // Initialize an empty map
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while decoding the API response"})
		logger.LogMessage("Request", "Error decoding the response: "+err.Error())
		return
	}

	var locationTagArray []string // Initialize an empty string array

	if items, ok := data["items"].([]interface{}); ok {
		for _, member := range items {
			if memberMap, ok := member.(map[string]interface{}); ok {
				if tag, ok := memberMap["tag"].(string); ok {
					locationTagArray = append(locationTagArray, tag)
				}
			}
		}
	}

	collectedTagArray := append(locationTagArray, clanTagsArray...) // Merge both arrays

	/* Remove duplicate tags and invalid tags */
	tagMap := make(map[string]bool)
	validTags := make([]string, 0)

	for _, tag := range collectedTagArray {
		// Check if tag starts with '#'
		if strings.HasPrefix(tag, "#") && !tagMap[tag] {
			validTags = append(validTags, tag)
			tagMap[tag] = true
		}
	}

	var result []interface{} // Initialize an empty array
	/* Process Validated Tags to get the fame history(result) */
	for _, getFameHistoryWithRiverLogCaller := range validTags {
		result = append(result, getFameHistoryWithRiverLog(getFameHistoryWithRiverLogCaller))
	}
	/* Print the end results of parameter processing */
	c.JSON(http.StatusOK, gin.H{"linechartRiverRaceLog": result}) // Return JSON object with "tags" key
}

// Get the fame history of a clan by Tag for the last 10 Weeks
func getFameHistoryWithRiverLog(tag string) []interface{} {
	urlForApiRequest := "https://api.clashroyale.com/v1/clans/" + url.QueryEscape(tag) + "/riverracelog"
	response, err := tools.MakeRequest(urlForApiRequest)
	if err != nil {
		logger.LogMessage("Request", "Error while making request: "+err.Error())
		return []interface{}{"ERROR: An error occurred while processing the request"}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.LogMessage("Request", "Error while closing response body: "+err.Error())
		}
	}(response.Body)

	var data struct {
		Items []struct {
			Standings []struct {
				Clan map[string]interface{} `json:"clan"`
			} `json:"standings"`
		} `json:"items"`
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		logger.LogMessage("Request", "Error while decoding response: "+err.Error())
		return []interface{}{"ERROR: An error occurred while decoding the response"}
	}

	/*------------Get Clan Name-----------*/
	getClanName := func(tag string) (string, error) {
		urlForApiRequest := "https://api.clashroyale.com/v1/clans/" + url.QueryEscape(tag)
		response, err := tools.MakeRequest(urlForApiRequest)
		if err != nil {
			logger.LogMessage("Request", "Error while making request: "+err.Error())
			return "", err
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				logger.LogMessage("Request", "Error while closing response body: "+err.Error())
			}
		}(response.Body)

		var clanData struct {
			Name string `json:"name"`
		}

		err = json.NewDecoder(response.Body).Decode(&clanData)
		if err != nil {
			logger.LogMessage("Request", "Error while decoding response: "+err.Error())
			return "", err
		}

		return clanData.Name, nil
	}

	clanName, err := getClanName(tag)
	if err != nil {
		return []interface{}{"ERROR: An error occurred while processing the request"}
	}

	/*------------Get Clan Fame History-----------*/
	var fameHistory []interface{}
	for index, item := range data.Items {
		for _, standing := range item.Standings {
			if standing.Clan["tag"] == tag {
				fameHistory = append(fameHistory, map[string]interface{}{
					"week": index,
					"fame": standing.Clan["fame"],
				})
			}
		}
	}

	/*------------Return Clan Name, Tag and Fame History-----------*/
	return []interface{}{
		map[string]interface{}{
			"clanName":        clanName,
			"clanTag":         tag,
			"clanFameHistory": fameHistory,
		},
	}
}
