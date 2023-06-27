package crt_person

import (
	"b4u/backend/logger"
)

func CalculateMember(members map[string]interface{}, tag string) (int, string, int) {
    memberData, ok := members["items"].([]interface{})
    if !ok {
        logger.LogMessage("Routine", "Error while retrieving member data.")
        return 0, "", 0
    }
    
    for _, member := range memberData {
        memberMap, ok := member.(map[string]interface{})
        if !ok {
            logger.LogMessage("Routine", "Error while extracting member data.")
            return 0, "", 0
        }
        
        memberTag, ok := memberMap["tag"].(string)
        if !ok {
            logger.LogMessage("Routine", "Error while extracting member tag.")
            return 0, "", 0
        }
        
        if memberTag == tag {
            clanRankFloat, ok := memberMap["clanRank"].(float64)
            if !ok {
                logger.LogMessage("Routine", "Error while extracting clanRank.")
                return 0, "", 0
            }
            clanRank := int(clanRankFloat)
            
            role, ok := memberMap["role"].(string)
            if !ok {
                logger.LogMessage("Routine", "Error while extracting role.")
                return 0, "", 0
            }
            
            trophiesFloat, ok := memberMap["trophies"].(float64)
            if !ok {
                logger.LogMessage("Routine", "Error while extracting trophies.")
                return 0, "", 0
            }
            trophies := int(trophiesFloat)
            
            return clanRank, role, trophies
        }
    }
    return 0, "", 0
}
