package tools

import (
	"b4u/backend/logger"
	"github.com/joho/godotenv"
	"net/url"
	"os"
)

var (
	// Request environment variables
	AccessToken    string
	ClanTag        string
	EncodedClanTag string
)

func LoadDotEnv() {
	// Load environment variables
	err := godotenv.Load() // Load environment variables from .env file
	if err != nil {        // If there was an error while loading the environment variables
		logger.LogMessage("Environment", "Error while loading environment variables: "+err.Error())
	}

	// Request environment variables
	AccessToken = os.Getenv("ACCESS_TOKEN")
	ClanTag = os.Getenv("CLAN_TAG")
	EncodedClanTag = url.QueryEscape(ClanTag)
}
