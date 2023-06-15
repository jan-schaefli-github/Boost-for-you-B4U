package art_clan

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"io"
	"net/url"
)

func getClanMembers(clanTag string) ([]byte, error) {
	// Construct the URL for the request
	apiUrl := "https://api.clashroyale.com/v1/clans/" + url.QueryEscape(clanTag) + "/members"

	// Send the request and get the response
	response, err := tools.MakeRequest(apiUrl)
	if err != nil {
		logger.LogMessage("Routine", "Error while making request: "+err.Error())
		return nil, err
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
		return nil, err
	}

	// Return the response body
	return body, nil
}
