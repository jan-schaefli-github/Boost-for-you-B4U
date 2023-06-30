package clanRegistry

import (
	"b4u/backend/tools"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TagData struct {
	Tags []string `yaml:"tags"`
}

// CreateRegister Register a clan and store tags in a YAML file
func CreateRegister(c *gin.Context) {
	clanTag := c.Query("clanTag")
	clanTag, err := url.QueryUnescape(clanTag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid clanTag"})
		return
	}

	// Validate clanTag
	if !strings.HasPrefix(clanTag, "#") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid clanTag. It should start with a '#' symbol."})
		return
	}

	// Store the clanTag in the YAML file
	err = storeClanTagInYAML(clanTag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store clanTag in YAML file", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Clan registered successfully and stored in YAML file"})
}

// Store the clanTag in a YAML file
func storeClanTagInYAML(clanTag string) error {
	filename := "clanRegistry/newRegistrys/tags.yaml"

	// Read existing YAML data
	yamlData, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read YAML file: %s", err.Error())
	}

	var data TagData
	err = yaml.Unmarshal(yamlData, &data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal YAML data: %s", err.Error())
	}

	// Check if clanTag already exists in the YAML data
	if isClanTagExists(clanTag, data.Tags) {
		return fmt.Errorf("clanTag already exists in YAML file")
	}

	// Initialize the map if it is nil
	if data.Tags == nil {
		data.Tags = make([]string, 0)
	}

	// Append clanTag to the list of tags
	data.Tags = append(data.Tags, clanTag)

	// Marshal the updated data back to YAML
	updatedYAMLData, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML data: %s", err.Error())
	}

	// Write the updated YAML data to the file
	err = ioutil.WriteFile(filename, updatedYAMLData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write YAML file: %s", err.Error())
	}

	return nil
}

// Check if clanTag already exists in the tags slice
func isClanTagExists(clanTag string, tags []string) bool {
	for _, tag := range tags {
		if tag == clanTag {
			return true
		}
	}
	return false
}

// WriteTagsToDatabase Write the stored tags from the YAML file to the database
// WriteTagsToDatabase Write the stored tags from the YAML file to the database
func WriteTagsToDatabase(c *gin.Context) {
	password := c.Query("password")

	// Validate the admin password
	if password != os.Getenv("ADMIN_PASS") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	// Connect to the database
	db, err := tools.ConnectToDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	// Read the stored tags from the YAML file
	tags, err := readTagsFromYAML("clanRegistry/newRegistrys/tags.yaml")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read tags from YAML file", "details": err.Error()})
		return
	}

	// Write the tags to the database
	err = writeTagsToDatabase(tags, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write tags to the database", "details": err.Error()})
		return
	}

	// Delete the tags from the YAML file
	err = deleteTagsFromYAML("clanRegistry/newRegistrys/tags.yaml")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tags from YAML file", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tags written to the database successfully"})
}

// Delete the tags from the YAML file
func deleteTagsFromYAML(filename string) error {
	emptyData := TagData{}

	// Marshal the empty data to YAML
	emptyYAMLData, err := yaml.Marshal(emptyData)
	if err != nil {
		return fmt.Errorf("failed to marshal empty YAML data: %s", err.Error())
	}

	// Write the empty YAML data to the file
	err = ioutil.WriteFile(filename, emptyYAMLData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write YAML file: %s", err.Error())
	}

	return nil
}

// Read the stored tags from the YAML file
func readTagsFromYAML(filename string) ([]string, error) {
	yamlData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML file: %s", err.Error())
	}

	var data TagData
	err = yaml.Unmarshal(yamlData, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML data: %s", err.Error())
	}

	return data.Tags, nil
}

// Write the tags to the database
func writeTagsToDatabase(tags []string, db *sql.DB) error {
	for _, tag := range tags {
		// Check if tag already exists in the database
		exists, err := checkClanExists(tag, db)
		if err != nil {
			return err
		}

		if exists {
			continue
		}

		// Insert the tag into the database
		err = InsertClanTag(tag, db)
		if err != nil {
			return err
		}
	}

	return nil
}

// Check if clanTag already exists in the database
func checkClanExists(clanTag string, db *sql.DB) (bool, error) {
	query := "SELECT COUNT(*) FROM clan WHERE tag = ?"
	var count int
	err := db.QueryRow(query, clanTag).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// InsertClanTag Insert the clanTag into the database
func InsertClanTag(clanTag string, db *sql.DB) error {
	query := "INSERT INTO clan (tag) VALUES (?)"
	_, err := db.Exec(query, clanTag)
	return err
}
