package main

import "fmt"

func processDayIdentifier(seasonIndex int, sectionIndex int, dayIndex int) string {
	identifier := fmt.Sprintf("%03d%03d%03d", seasonIndex, sectionIndex, dayIndex)
	return identifier
}
