package art_clan

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"encoding/json"
	"io"
	"net/url"
)

func GetRiverracelog(clanTag string) map[string]interface{} {

	// Construct the URL for the request
	apiUrl := "https://api.clashroyale.com/v1/clans/" + url.QueryEscape(clanTag) + "/riverracelog"

	// Send the request and get the response
	response, err := tools.MakeRequest(apiUrl)
	if err != nil {
		logger.LogMessage("Routine", "Error while making request: "+err.Error())
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.LogMessage("Routine", "Error while closing response body: "+err.Error())
		}
	}(response.Body)

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.LogMessage("Routine", "Error while reading response: "+err.Error())
		return nil
	}

	// Unmarshal the response body into a map[string]interface{}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		logger.LogMessage("Routine", "Error while unmarshaling response: "+err.Error())
		return nil
	}

	return data
}
