package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
)

// Routine to collect data from the API
func dataCollector() {

	var loopedPlayers []string

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

		// Check if the person is already in the database
		if !checkPerson(tag) {
			clanTag := getClanTag(tag)

			// Check if the clan tag is empty
			if clanTag == "" {
				logMessage("Routine", "Clan tag is empty: "+tag)
			} else {
				createPerson(tag, clanTag)
			}
		} else {
			clanTag := getClanTag(tag)

			// Check if the clan tag is empty
			if clanTag == "" {
				logMessage("Routine", "Clan tag is empty: "+tag)
			} else {
				loopedPlayers = append(loopedPlayers, tag)

				// Save participant data
				err = saveParticipantData(tag, fame, decksUsedToday)
				if err != nil {
					logMessage("Routine", "Error while saving participant data: "+err.Error())
					return
				}
			}
		}
	}

	// Update person status
	updatePersonStatus(loopedPlayers)
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

// Routine to update the status of the persons
func updatePersonStatus(loopedPlayers []string) {

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
		"ELSE 0 END"

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

// ------------------------------ TEMPORARY ------------------------------
// Routine to save the participant data to the database
func saveParticipantData(tag string, fame, decksUsedToday float64) error {

	filePath := "database.txt"

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logMessage("Routine", "Error while closing file: "+err.Error())
			return
		}
	}(file)

	data := fmt.Sprintf("Tag: %s, Fame: %.0f, DecksUsedToday: %.0f\n", tag, fame, decksUsedToday)

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}
