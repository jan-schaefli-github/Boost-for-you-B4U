package rt_main

import (
	"b4u/backend/routine/v1/rt_api/art_clan"
	"b4u/backend/routine/v1/rt_calculations/crt_clan"
	"b4u/backend/routine/v1/rt_database/drt_clan"
)

func Routine() {

	clanTags := drt_clan.GetClanTags()

	for _, clanTag := range clanTags {
		currentriverrace := art_clan.GetCurrentriverrace(clanTag)

		periodIndex, periodType, sectionIndex := crt_clan.CalculateCurrentPeriod(currentriverrace)

		if periodType == "warDay" {

		} else {

		}

		println(periodIndex)
		println(periodType)
		println(sectionIndex)
		println(clanTag)
	}
}
