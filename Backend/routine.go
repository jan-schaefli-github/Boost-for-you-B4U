package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

// Routine to collect data from the API
func dataCollector(clanTags []string) {

	for _, clanTag := range clanTags {

		encodedClanTag := url.QueryEscape(clanTag)

		warStatus, periodIndex := checkWarStatus(clanTag)

		if warStatus {

			// Get clan members
			members := getClanMembers(encodedClanTag)

			updatePersonStatus(members, clanTag)

			for _, member := range members {
				if !checkPerson(member) {
					createPerson(member, clanTag)
				}
			}

			// Get clan data
			generalurl := "https://api.clashroyale.com/v1/clans/" + encodedClanTag + "/currentriverrace"
			response, err := makeRequest(generalurl)
			if err != nil {
				logMessage("Routine", "Error while making request: "+err.Error())
				return
			}

			// Close response body
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					logMessage("Routine", "Error while closing response body: "+err.Error())
					return
				}
			}(response.Body)

			// Decode response
			var data map[string]interface{}
			err = json.NewDecoder(response.Body).Decode(&data)
			if err != nil {
				logMessage("Routine", "Error while decoding response: "+err.Error())
				return
			}

			// Check if there is any clan data
			clan, ok := data["clan"].(map[string]interface{})
			if !ok {
				logMessage("Routine", "Error while extracting clan data from response.")
				return
			}

			// Check if there is any participants data
			participants, ok := clan["participants"].([]interface{})
			if !ok {
				logMessage("Routine", "Error while extracting participants from clan data.")
				return
			}

			// Check if there are any participants in the clan
			if len(participants) == 0 {
				logMessage("Routine", "No participants found.")
				return
			}

			// Loop through participants
			for _, participant := range participants {
				participantData, ok := participant.(map[string]interface{})
				if !ok {
					logMessage("Routine", "Error while extracting participant data from participants.")
					return
				}

				// Extract tag from participant data
				tag, ok := participantData["tag"].(string)
				if !ok {
					logMessage("Routine", "Error while extracting tag from participant data.")
					return
				}

				// Extract fame from participant data
				fame, ok := participantData["fame"].(float64)
				if !ok {
					logMessage("Routine", "Error while extracting fame from participant data.")
					return
				}

				// Extract decksUsedToday from participant data
				decksUsedToday, ok := participantData["decksUsedToday"].(float64)
				if !ok {
					logMessage("Routine", "Error while extracting decksUsedToday from participant data.")
					return
				}

				clanTag := getClanTag(tag)

				// Check if the clan tag is empty
				if clanTag == "" {
					logMessage("Routine", "Clan tag is empty: "+tag)
				} else {

					// Save participant data
					if checkPerson(tag) {
						createDailyReport(decksUsedToday, fame, periodIndex, tag)
					} else {
						logMessage("Routine", "Person does not exist in database: "+tag)
					}
				}
			}
		} else {
			logMessage("Routine", "War is not active for clan: "+clanTag)
		}
	}
}

// Routine to check if the person is already in the database
func getClanTag(playerTag string) string {

	encodedPlayerTag := url.QueryEscape(playerTag)

	// Get player data
	generalurl := "https://api.clashroyale.com/v1/players/" + encodedPlayerTag
	response, err := makeRequest(generalurl)
	if err != nil {
		logMessage("Routine", "Error while making request: "+err.Error())
		return ""
	}

	// Close response body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logMessage("Routine", "Error while closing response body: "+err.Error())
			return
		}
	}(response.Body)

	// Decode response
	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		logMessage("Routine", "Error while decoding response: "+err.Error())
		return ""
	}

	// Extract clan data from response
	clan, ok := data["clan"].(map[string]interface{})

	// Check if there is any clan data
	if clan == nil {
		return ""
	}

	// Check if there is any clan tag
	if !ok {
		logMessage("Routine", "Error while extracting clan data from response.")
		return ""
	}

	// Extract tag from clan data
	tag, ok := clan["tag"].(string)
	if !ok {
		logMessage("Routine", "Error while extracting tag from clan data.")
		return ""
	}

	return tag
}

// Routine to check if the person is still in the clan
func getClanMembers(encodedClanTag string) []string {

	// Construct the URL for the request
	url := "https://api.clashroyale.com/v1/clans/" + encodedClanTag + "/members"

	// Send the request and get the response
	response, err := makeRequest(url)
	if err != nil {
		logMessage("Routine", "Error while making request: "+err.Error())
		return nil
	}
	defer response.Body.Close()

	// Decode the response
	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		logMessage("Routine", "Error while decoding response: "+err.Error())
		return nil
	}

	// Extract the members from the response
	members, ok := data["items"].([]interface{})
	if !ok {
		logMessage("Routine", "Error while extracting members from response.")
		return nil
	}

	var memberTags []string

	// Loop through the members and extract the tags
	for _, member := range members {

		// Extract the member data
		memberMap, ok := member.(map[string]interface{})
		if !ok {
			logMessage("Routine", "Error while extracting member data from members.")
			return nil
		}

		// Extract the tag from the member data
		tag, ok := memberMap["tag"].(string)
		if !ok {
			logMessage("Routine", "Error while extracting tag from member data.")
			return nil
		}

		memberTags = append(memberTags, tag)
	}

	return memberTags
}

// Routine to update the status of the persons
func updatePersonStatus(loopedPlayers []string, clanTag string) {

	// Connect to database
	db, err := connectToDatabase()
	if err != nil {
		logMessage("Routine", "Error while connecting to database: "+err.Error())
		return
	}

	// Close database connection
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logMessage("Routine", "Error while closing database: "+err.Error())
			return
		}
	}(db)

	// Create a string of placeholders for the query
	placeholders := strings.Repeat("?, ", len(loopedPlayers)-1) + "?"

	// Create the query
	query := "UPDATE person SET clanStatus = CASE " +
		"WHEN tag IN (" + placeholders + ") THEN 1 " +
		"ELSE 0 END" +
		" WHERE fk_clan = '" + clanTag + "'"

	// Prepare the query
	stmt, err := db.Prepare(query)
	if err != nil {
		logMessage("Routine", "Error while preparing query: "+err.Error())
		return
	}

	// Close the statement
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			logMessage("Routine", "Error while closing statement: "+err.Error())
			return
		}
	}(stmt)

	// Create a slice of interface{} with the length of loopedPlayers
	args := make([]interface{}, len(loopedPlayers))
	for i, player := range loopedPlayers {
		args[i] = player
	}

	// Execute the query
	_, err = stmt.Exec(args...)
	if err != nil {
		logMessage("Routine", "Error while executing query: "+err.Error())
		return
	}

	return
}

// Routine to check if the clan is in war
func checkWarStatus(clanTag string) (bool, int) {

	// Construct the URL for the request
	url := "https://api.clashroyale.com/v1/clans/" + url.QueryEscape(clanTag) + "/currentriverrace"

	// Send the request and get the response
	response, err := makeRequest(url)
	if err != nil {
		logMessage("Routine", "Error while making request: "+err.Error())
		return false, 0
	}
	defer response.Body.Close()

	// Decode the response into a map object
	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		logMessage("Routine", "Error while decoding response: "+err.Error())
		return false, 0
	}

	// Extract the war status from the response
	periodType, ok := data["periodType"].(string)
	if !ok {
		logMessage("Routine", "Error while extracting periodType from response.")
		return false, 0
	}

	periodIndex, ok := data["periodIndex"].(float64)
	if !ok {
		logMessage("Routine", "Error while extracting periodIndex from response.")
		return false, 0
	}

	// Check if the clan is in war
	warStatus := false
	if periodType == "warDay" {
		warStatus = true
	}

	return warStatus, int(periodIndex)
}
