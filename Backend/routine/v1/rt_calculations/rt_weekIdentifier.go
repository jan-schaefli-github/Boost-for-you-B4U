package rt_calculations

import (
	"fmt"
)

func CalculateWeekIdentifier(seasonId int, sectionIndex int) string {
	identifier := fmt.Sprintf("%03d%03d", seasonId, sectionIndex)

	return identifier
}