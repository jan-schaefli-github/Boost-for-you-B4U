package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func dataCollector() {
	url := "https://api.clashroyale.com/v1/clans/" + encodedClanTag + "/currentriverrace"
	response, err := makeRequest(url)
	if err != nil {
		log.Fatal("Fehler beim Abrufen der aktuellen Riverrace-Informationen:", err)
	}
	defer response.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal("Fehler beim Dekodieren der API-Antwort:", err)
	}

	clan, ok := data["clan"].(map[string]interface{})
	if !ok {
		log.Fatal("Fehler beim Extrahieren des Clan-Datenobjekts aus der API-Antwort.")
	}

	participants, ok := clan["participants"].([]interface{})
	if !ok {
		log.Fatal("Fehler beim Extrahieren der Teilnehmerliste aus der Clan-Datenstruktur.")
	}

	if len(participants) == 0 {
		log.Fatal("Die Teilnehmerliste ist leer.")
	}

	for _, participant := range participants {
		participantData, ok := participant.(map[string]interface{})
		if !ok {
			log.Fatal("Fehler beim Extrahieren des Teilnehmer-Datenobjekts aus der Teilnehmerliste.")
		}

		tag, ok := participantData["tag"].(string)
		if !ok {
			log.Fatal("Fehler beim Extrahieren des Clan-Tags aus dem Teilnehmer-Datenobjekt.")
		}

		fame, ok := participantData["fame"].(float64)
		if !ok {
			log.Fatal("Fehler beim Extrahieren des Ruhms aus dem Teilnehmer-Datenobjekt.")
		}

		decksUsedToday, ok := participantData["decksUsedToday"].(float64)
		if !ok {
			log.Fatal("Fehler beim Extrahieren der Anzahl der heute verwendeten Decks aus dem Teilnehmer-Datenobjekt.")
		}

		err = saveParticipantData(tag, fame, decksUsedToday)
		if err != nil {
			log.Fatal("Fehler beim Speichern der Teilnehmerdaten:", err)
		}
	}
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
