package rt_calculations

import (
	"fmt"
)

func CalculateWeekIdentifier(sectionIndex int) string {
	identifier := fmt.Sprintf("%03d", sectionIndex)

	return identifier
}