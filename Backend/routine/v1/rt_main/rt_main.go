package rt_main

import (
	"b4u/backend/logger"
	"b4u/backend/routine/v1/rt_api/art_clan"
	"b4u/backend/routine/v1/rt_api/art_person"
	"b4u/backend/routine/v1/rt_calculations"
	"b4u/backend/routine/v1/rt_calculations/crt_clan"
	"b4u/backend/routine/v1/rt_calculations/crt_person"
	"b4u/backend/routine/v1/rt_database/drt_clan"
	"b4u/backend/routine/v1/rt_database/drt_person"
)

func Routine() {

	// Get the clan tags
	clanTags := drt_clan.GetClanTags()

	// For each clan tag
	for _, clanTag := range clanTags {

		// Get the clan members
		members := art_clan.GetMembers(clanTag)

		// Get war information
		currentriverrace := art_clan.GetCurrentriverrace(clanTag)
		periodIndex, periodType, sectionIndex := crt_clan.CalculateCurrentPeriod(currentriverrace)

		// Get the participants
		participants := crt_person.CalculateParticipants(currentriverrace)

		// For each participant
		for _, participant := range participants {

			participantData, ok := participant.(map[string]interface{})
			if !ok {
				logger.LogMessage("Routine", "Error while extracting participant from participants.")
				return
			}

			// Calculate the participant data
			tag, name, fame, repairPoints, boatAttacks, decksUsedToday := crt_person.CalculateParticipantData(participantData)

			clanRank, role, trophies := crt_person.CalculateMember(members, tag)

			// Calculate the day identifier
			dayIdentifier := rt_calculations.CalculateDayIdentifier(sectionIndex, periodIndex)
			decksUsedYesterday, lastFame, lastDayIdentifier := drt_person.GetLastDailyReport(tag)

			// Calculate the week identifier
			weekIdentifier := rt_calculations.CalculateWeekIdentifier(sectionIndex)
			lastmissedDecks, lastWeekIdentifier := drt_person.GetLastWeeklyReport(tag)

			// Get the person clan tag
			personData := art_person.GetPerson(tag)
			personClanTag := crt_person.CalculatePersonClan(personData)

			// Check if the person is in the clan
			if clanTag == personClanTag {

				// Check if Person already exists in the database
				if !drt_person.CheckPerson(tag) {

					drt_person.CreatePerson(tag, name, personClanTag)
				}

				// Update the daily report
				if lastDayIdentifier == dayIdentifier {

					drt_person.UpdateDailyReport(decksUsedToday, fame, repairPoints, boatAttacks, dayIdentifier, tag)
				} else {

					// If it is war day
					if periodType == "warDay" {

						// Calculate the missed decks
						missedDecks := lastmissedDecks + (4 - decksUsedYesterday)

						// Update the weekly report
						drt_person.UpdateWeeklyReport(lastFame, missedDecks, weekIdentifier, tag)

					} else {
						// Update the weekly report
						drt_person.UpdateWeeklyReport(0, 0, weekIdentifier, tag)
					}

					// Check if the week changed
					if lastWeekIdentifier != weekIdentifier {

						// NEW WEEK STARTED
						drt_person.CreateWeeklyReport(weekIdentifier, tag)
					}

					// NEW DAY STARTED
					drt_person.CreateDailyReport(decksUsedToday, fame, repairPoints, boatAttacks, dayIdentifier, tag)
				}

				// Update the person
				drt_person.UpdatePerson(tag, role, trophies, clanRank)
			}
		}

		// Update status
		memberTags := crt_clan.CalculateMemberTags(members)

		drt_person.UpdatePersonStatus(memberTags, clanTag)
	}
	logger.LogMessage("Routine", "Routine finished.")
}
