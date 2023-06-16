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
