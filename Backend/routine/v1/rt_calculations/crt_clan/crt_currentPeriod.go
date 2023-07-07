package crt_clan

import (
	"b4u/backend/logger"
)

func CalculateCurrentPeriod(data map[string]interface{}) (int, string, int) {

	// Extract the period index from the response
	periodIndexFloat, ok := data["periodIndex"].(float64)
	if !ok {
		logger.LogMessage("Routine", "Error while extracting periodIndex from response.")
		return 0, "", 0
	}
	periodIndex := int(periodIndexFloat)

	// Extract the period type from the response
	periodType, ok := data["periodType"].(string)
	if !ok {
		logger.LogMessage("Routine", "Error while extracting periodType from response.")
		return 0, "", 0
	}

	// Extract the section index from the response
	sectionIndexFloat, ok := data["sectionIndex"].(float64)
	if !ok {
		logger.LogMessage("Routine", "Error while extracting sectionIndex from response.")
		return 0, "", 0
	}
	sectionIndex := int(sectionIndexFloat)

	return periodIndex, periodType, sectionIndex
}

func CalculateCurrentPeriodSeason(data map[string]interface{}, sectionIndex int) (int) {
	
	// Extract the items array from the response
	items, ok := data["items"].([]interface{})
	if !ok {
		return 0
	}

	seasonId, ok := items[0].(map[string]interface{})["seasonId"].(float64)
	if !ok {
		return 0
	}

	if sectionIndex == 0 {
		seasonId = seasonId + 1
	}

	return int(seasonId)
}
