package clanRegistry

import (
	"b4u/backend/logger"
	"b4u/backend/tools"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Possible YAML 2 step registration.

/*type Person struct {
	Name  string `yaml:"name"`
	Age   int    `yaml:"age"`
	Email string `yaml:"email"`
}

func GetClanRegistry() {
	// Specify the folder path
	folderPath := "clanRegistry/newRegistrys"

	// Open the folder
	folder, err := os.Open(folderPath)
	if err != nil {
		logger.LogMessage("Routine", "No newRegistrys Folder found")
		return
	}
	defer func(folder *os.File) {
		err := folder.Close()
		if err != nil {
			logger.LogMessage("Routine", "Something went wrong closing the folder")
		}
	}(folder)

	// Get the file names in the folder
	fileNames, err := folder.Readdirnames(-1)
	if err != nil {
		panic(err)
	}

	// Create a slice to hold the persons
	var persons []Person

	// Iterate through the file names
	for _, fileName := range fileNames {
		// Check if the file has a YAML extension
		if strings.HasSuffix(fileName, ".yaml") || strings.HasSuffix(fileName, ".yml") {
			// Open the YAML file
			filePath := filepath.Join(folderPath, fileName)
			file, err := os.Open(filePath)
			if err != nil {
				panic(err)
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					logger.LogMessage("Routine", fmt.Sprintf("Something went wrong closing the file: %s", filePath))
				}
			}(file)

			// Read the YAML data from the file
			yamlData, err := io.ReadAll(file)
			if err != nil {
				panic(err)
			}

			// Unmarshal the YAML data into the persons slice
			var yamlPersons []Person
			err = yaml.Unmarshal(yamlData, &yamlPersons)
			if err != nil {
				panic(err)
			}

			// Append the YAML persons to the main persons slice
			persons = append(persons, yamlPersons...)
		}
	}

	// Print the persons
	for _, person := range persons {
		fmt.Printf("Name: %s\nAge: %d\nEmail: %s\n\n", person.Name, person.Age, person.Email)
	}
}*/

func CreateRegister(c *gin.Context) {
	clanTag := c.Query("clanTag")
	clanTag, err := url.QueryUnescape(clanTag)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid clanTag"})
		return
	}
	fmt.Println(clanTag)

	if !strings.HasPrefix(clanTag, "#") {
		// Return an error response if clanTag does not start with "#"
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid clanTag. It should start with a '#' symbol."})
		return
	}

	db, err := tools.ConnectToDatabase()
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	existingTags, err := db.Query("SELECT tag FROM clan")
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query the database", "details": err.Error()})
		return
	}

	defer func(existingTags *sql.Rows) {
		err := existingTags.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data from the database", "details": err.Error()})
			return
		}
	}(existingTags)

	for existingTags.Next() {
		var tag string
		err := existingTags.Scan(&tag)
		if err != nil {
			// Handle the error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data from the database", "details": err.Error()})
			return
		}

		if tag == clanTag {
			// Tag already exists in the database
			if !checkTheClan(tag) {
				c.JSON(http.StatusOK, gin.H{"message": "Clan doesn't exist"})
				return
			} else {
				// Clan exists, return success response
				c.JSON(http.StatusOK, gin.H{"message": "Clan already exists"})
				return
			}
		}
	}

	// Tag does not exist in the database, insert the tag
	err = InsertClanTag(clanTag)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert tag into the database", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Clan registered successfully"})
}

func InsertClanTag(tag string) error {
	db, err := tools.ConnectToDatabase()
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %s", err.Error())
	}

	query := "INSERT INTO clan (tag) VALUES (?)"
	_, err = db.Exec(query, tag)
	if err != nil {
		return fmt.Errorf("failed to insert tag into the database: %s", err.Error())
	}

	return nil
}

func checkTheClan(tag string) bool {
	urlForApiRequest := "https://api.clashroyale.com/v1/clans/" + url.QueryEscape(tag)
	response, err := tools.MakeRequest(urlForApiRequest)
	if err != nil {
		logger.LogMessage("Request", "Error while making request: "+err.Error())
		return false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.LogMessage("Request", "Error while closing response body: "+err.Error())
		}
	}(response.Body)

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		logger.LogMessage("Request", "Error while decoding response: "+err.Error())
		return false
	}

	reason, found := data["reason"]
	if found && reason == "notFound" {
		return false
	}

	return true
}
