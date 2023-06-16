package crt_person

import (
	"b4u/backend/logger"
)

func CalculatePersonClan(data map[string]interface{}) (string) {

	// Extract the clan from the response
	clan, ok := data["clan"].(map[string]interface{})
	if !ok {
		//logger.LogMessage("Routine", "Error while extracting clan from response.")
		return ""
	}

	tag, ok := clan["tag"].(string)
	if !ok {
		logger.LogMessage("Routine", "Error while extracting tag from clan.")
		return ""
	}

	return tag
}