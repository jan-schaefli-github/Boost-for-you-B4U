package rt_calculations

import (
	"fmt"
)

func CalculateDayIdentifier(seasonId int, sectionIndex int, periodIndex int) string {
	identifier := fmt.Sprintf("%03d%03d%03d", seasonId, sectionIndex, periodIndex)

	return identifier
}