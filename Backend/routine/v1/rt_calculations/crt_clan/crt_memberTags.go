package crt_clan

import (
	"b4u/backend/logger"
)

func CalculateMemberTags(members map[string]interface{}) []string {
	var tags []string

	items, ok := members["items"].([]interface{})
	if !ok {
		logger.LogMessage("Routine", "Error while extracting items from members.")
		return tags
	}

	for _, item := range items {
		member, ok := item.(map[string]interface{})
		if !ok {
			logger.LogMessage("Routine", "Error while extracting member from items.")
			continue
		}

		tag, ok := member["tag"].(string)
		if ok {
			tags = append(tags, tag)
		} else if !ok {
			logger.LogMessage("Routine", "Error while extracting tag from member.")
			continue
		}
	}

	return tags
}
