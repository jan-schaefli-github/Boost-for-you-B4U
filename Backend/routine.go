package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
)

// Routine to collect data from the API
func dataCollector() {

	loopedPlayers := []string{}

	url := "https://api.clashroyale.com/v1/clans/" + encodedClanTag + "/currentriverrace"
	response, err := makeRequest(url)
	if err != nil {
		logMessage("Routine", "Error while making request: "+err.Error())
		return
	}
	defer response.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		logMessage("Routine", "Error while decoding response: "+err.Error())
		return
	}

	clan, ok := data["clan"].(map[string]interface{})
	if !ok {
		logMessage("Routine", "Error while extracting clan data from response.")
		return
	}

	participants, ok := clan["participants"].([]interface{})
	if !ok {
		logMessage("Routine", "Error while extracting participants from clan data.")
		return
	}

	// Check if there are any participants
	if len(participants) == 0 {
		logMessage("Routine", "No participants found.")
		return
	}

	for _, participant := range participants {
		participantData, ok := participant.(map[string]interface{})
		if !ok {
			logMessage("Routine", "Error while extracting participant data from participants.")
			return
		}

		tag, ok := participantData["tag"].(string)
		if !ok {
			logMessage("Routine", "Error while extracting tag from participant data.")
			return
		}

		fame, ok := participantData["fame"].(float64)
		if !ok {
			logMessage("Routine", "Error while extracting fame from participant data.")
			return
		}

		decksUsedToday, ok := participantData["decksUsedToday"].(float64)
		if !ok {
			logMessage("Routine", "Error while extracting decksUsedToday from participant data.")
			return
		}

		if !checkPerson(tag) {
			clanTag := getClanTag(tag)
			if clanTag == "" {
				logMessage("Routine", "Clan tag is empty: "+tag)
			} else {
				createPerson(tag, clanTag)
			}
		} else {
			clanTag := getClanTag(tag)
			if clanTag == "" {
				logMessage("Routine", "Clan tag is empty: "+tag)
			} else {
				loopedPlayers = append(loopedPlayers, tag)
				err = saveParticipantData(tag, fame, decksUsedToday)
				if err != nil {
					logMessage("Routine", "Error while saving participant data: "+err.Error())
				}
			}
		}
	}
	updatePersonStatus(loopedPlayers)
}

func getClanTag(playerTag string) string {

	encodedPlayerTag := url.QueryEscape(playerTag)
	url := "https://api.clashroyale.com/v1/players/" + encodedPlayerTag
	response, err := makeRequest(url)
	if err != nil {
		logMessage("Routine", "Error while making request: "+err.Error())
	}
	defer response.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		logMessage("Routine", "Error while decoding response: "+err.Error())
	}

	clan, ok := data["clan"].(map[string]interface{})

	if clan == nil {
		return ""
	}

	if !ok {
		logMessage("Routine", "Error while extracting clan data from response.")
		return ""
	}

	tag, ok := clan["tag"].(string)
	if !ok {
		logMessage("Routine", "Error while extracting tag from clan data.")
		return ""
	}

	return tag
}

func saveParticipantData(tag string, fame, decksUsedToday float64) error {

	filePath := "database.txt"

	// Open the file in append mode and create it if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Prepare the data string to be written to the file
	data := fmt.Sprintf("Tag: %s, Fame: %.0f, DecksUsedToday: %.0f\n", tag, fame, decksUsedToday)

	// Write the data to the file
	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

// Routine to update the status of the persons
func updatePersonStatus(loopedPlayers []string) {
	db, err := connectToDatabase()
	if err != nil {
		logMessage("Routine", "Error while connecting to database: "+err.Error())
		return
	}
	defer db.Close()

	// Generate placeholders for the IN clause based on the number of loopedPlayers
	placeholders := strings.Repeat("?, ", len(loopedPlayers)-1) + "?"

	// Build the query with the placeholders and loopedPlayers values
	query := "UPDATE person SET clanStatus = CASE " +
		"WHEN tag IN (" + placeholders + ") THEN 1 " +
		"ELSE 0 END"

	// Prepare the query statement
	stmt, err := db.Prepare(query)
	if err != nil {
		logMessage("Routine", "Error while preparing query: "+err.Error())
		return
	}
	defer stmt.Close()

	// Create a slice of interface{} to hold the loopedPlayers values
	args := make([]interface{}, len(loopedPlayers))
	for i, player := range loopedPlayers {
		args[i] = player
	}

	// Execute the query with the loopedPlayers values
	_, err = stmt.Exec(args...)
	if err != nil {
		logMessage("Routine", "Error while executing query: "+err.Error())
		return
	}

	return
}
