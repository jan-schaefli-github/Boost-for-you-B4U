package rt_calculations

import (
	"fmt"
)

func CalculateDayIdentifier(sectionIndex int, periodIndex int) string {
	identifier := fmt.Sprintf("%03d%03d", sectionIndex, periodIndex)

	return identifier
}