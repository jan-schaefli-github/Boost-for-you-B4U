package crt_person

import (
	"b4u/backend/logger"
)

func CalculateParticipants(data map[string]interface{}) ([]interface{}) {

	// Extract the clan from the response
	clan, ok := data["clan"].(map[string]interface{})
	if !ok {
		logger.LogMessage("Routine", "Error while extracting clan from response.")
		return nil
	}

	participants, ok := clan["participants"].([]interface{})
	if !ok {
		logger.LogMessage("Routine", "Error while extracting participants from response.")
		return nil
	}

	return participants
}

func CalculateParticipantData(participant map[string]interface{}) (string, string, float64, float64, float64, float64) {

	// Extract the tag from the participant
	tag, ok := participant["tag"].(string)
	if !ok {
		logger.LogMessage("Routine", "Error while extracting tag from participant.")
		return "", "", 0, 0, 0, 0
	}

	name, ok := participant["name"].(string)
	if !ok {
		logger.LogMessage("Routine", "Error while extracting name from participant.")
		return "", "", 0, 0, 0, 0
	}

	// Extract the fame from the participant
	fame, ok := participant["fame"].(float64)
	if !ok {
		logger.LogMessage("Routine", "Error while extracting fame from participant.")
		return "", "", 0, 0, 0, 0
	}

	// Extract the repairPoints from the participant
	repairPoints, ok := participant["repairPoints"].(float64)
	if !ok {
		logger.LogMessage("Routine", "Error while extracting repairPoints from participant.")
		return "", "", 0, 0, 0, 0
	}

	// Extract the boatAttacks from the participant
	boatAttacks, ok := participant["boatAttacks"].(float64)
	if !ok {
		logger.LogMessage("Routine", "Error while extracting boatAttacks from participant.")
		return "", "", 0, 0, 0, 0
	}

	// Extract the decksUsedToday from the participant
	decksUsedToday, ok := participant["decksUsedToday"].(float64)
	if !ok {
		logger.LogMessage("Routine", "Error while extracting decksUsedToday from participant.")
		return "", "", 0, 0, 0, 0
	}

	return tag, name, fame, repairPoints, boatAttacks,  decksUsedToday
}