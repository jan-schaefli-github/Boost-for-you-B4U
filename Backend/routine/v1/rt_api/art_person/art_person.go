package art_person

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"io/ioutil"
	"net/url"
	"encoding/json"
)

func GetPerson(personTag string) (map[string]interface{}) {

	// Construct the URL for the request
	url := "https://api.clashroyale.com/v1/players/" + url.QueryEscape(personTag)

	// Send the request and get the response
	response, err := tools.MakeRequest(url)
	if err != nil {
		logger.LogMessage("Routine", "Error while making request: " + err.Error())
		return nil
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.LogMessage("Routine", "Error while reading response: " + err.Error())
		return nil
	}

	// Unmarshal the response body into a map[string]interface{}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		logger.LogMessage("Routine", "Error while unmarshaling response: " + err.Error())
		return nil
	}

	return data
}