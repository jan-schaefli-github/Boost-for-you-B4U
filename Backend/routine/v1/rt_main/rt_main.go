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
		riverracelog := art_clan.GetRiverracelog(clanTag)

		// Calculate the current period
		periodIndex, periodType, sectionIndex := crt_clan.CalculateCurrentPeriod(currentriverrace)
		seasonId := crt_clan.CalculateCurrentPeriodSeason(riverracelog, sectionIndex)

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
			dayIdentifier := rt_calculations.CalculateDayIdentifier(seasonId, sectionIndex, periodIndex)

			// Calculate the week identifier
			weekIdentifier := rt_calculations.CalculateWeekIdentifier(seasonId, sectionIndex)


			// Get last day data
			fameYesterday, decksUsedYesterday, missedDecksYesterday, repairPointsYesterday, boatAttacksYesterday := drt_person.GetLastDailyReport(tag, dayIdentifier)

			lastDayIdentifier := drt_person.GetLastDayIdentifier(tag)

			// Get last week data
			lastFame, lastDecksUsed, lastMissedDecks, lastRepairPoints, lastBoatAttacks, lastWeekIdentifier := drt_person.GetLastWeeklyReport(tag)

			// Get the person clan tag
			personData := art_person.GetPerson(tag)
			personClanTag := crt_person.CalculatePersonClan(personData)

			// Check if the person is in the clan
			if clanTag == personClanTag {

				// Check if Person already exists in the database
				if !drt_person.CheckPerson(tag) {

					drt_person.CreatePerson(tag, name, role, trophies, clanRank, clanTag)
				} else {

					// Check if the day changed
					if lastDayIdentifier != dayIdentifier {

						// Get last person data
						lastWholeFame, lastWholeDecksUsed, lastWholeMissedDecks, lastWholeRepairPoints, lastWholeBoatAttacks := drt_person.GetPerson(tag)

						// Calculate the whole report
						wholeFame := int(lastWholeFame) + int(fameYesterday)
						wholeDecksUsed := int(lastWholeDecksUsed) + int(decksUsedYesterday)
						wholeMissedDecks := int(lastWholeMissedDecks) + int(missedDecksYesterday)
						wholeRepairPoints := int(lastWholeRepairPoints) + int(repairPointsYesterday)
						wholeBoatAttacks := int(lastWholeBoatAttacks) + int(boatAttacksYesterday)

						// Update the person
						drt_person.UpdatePerson(tag, name, role, trophies, clanRank, wholeFame, wholeDecksUsed, wholeMissedDecks, wholeRepairPoints, wholeBoatAttacks)
					}
				}
				// Calculate the daily report
				fameToday := int(fame) - int(lastFame)
				missedDecksToday := 0
				repairPointsToday := int(repairPoints) - int(lastRepairPoints)
				boatAttacksToday := int(boatAttacks) - int(lastBoatAttacks)

				// Calculate the weekly report
				fameThisWeek := int(lastFame) + int(fameYesterday)
				decksUsedThisWeek := int(lastDecksUsed) + int(decksUsedYesterday)
				missedDecksThisWeek := int(lastMissedDecks) + int(missedDecksYesterday)
				repairPointsThisWeek := int(lastRepairPoints) + int(repairPointsYesterday)
				boatAttacksThisWeek := int(lastBoatAttacks) + int(boatAttacksYesterday)


				// Check if the period is not training
				if periodType != "training" {

					// Calculate the missed decks today
					missedDecksToday = 4 - int(decksUsedToday)

					// Calculate the missed decks this week
					missedDecksThisWeek = lastMissedDecks + missedDecksYesterday
				}

				// Update the daily report
				if lastDayIdentifier == dayIdentifier {

					drt_person.UpdateDailyReport(fameToday, int(decksUsedToday), missedDecksToday, repairPointsToday, boatAttacksToday, dayIdentifier, tag)
				} else {
					
						// Update the weekly report
						drt_person.UpdateWeeklyReport(fameThisWeek, decksUsedThisWeek, missedDecksThisWeek, repairPointsThisWeek, boatAttacksThisWeek, lastWeekIdentifier, tag)

					// Check if the week changed
					if lastWeekIdentifier != weekIdentifier {

						// NEW WEEK STARTED
						drt_person.CreateWeeklyReport(weekIdentifier, tag)

						// NEW DAY STARTED
						drt_person.CreateDailyReport(int(fame), int(decksUsedToday), int(missedDecksToday), int(repairPoints), int(boatAttacks), dayIdentifier, tag)
					} else {

						// NEW DAY STARTED
						drt_person.CreateDailyReport(fameToday, int(decksUsedToday), missedDecksToday, repairPointsToday, boatAttacksToday, dayIdentifier, tag)
					}
				}
			}
		}

		// Update status every hour
		memberTags := crt_clan.CalculateMemberTags(members)

		drt_person.UpdatePersonStatus(memberTags, clanTag)
	}
	logger.LogMessage("Routine", "Routine finished.")
}
