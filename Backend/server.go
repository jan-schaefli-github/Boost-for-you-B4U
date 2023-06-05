package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	// Request environment variables
	accessToken    string
	clanTag        string
	encodedClanTag string

	// Database environment variables
	username string
	password string
	host     string
	port     int
)

// -------------------- Start Structs -------------------- //

/*type Person struct {
	tag        string
	wholeFame  string
	clanStatus string
	joinDate   string
	fk_clan    string
}*/

type Person struct {
	Tag        string `json:"tag"`
	WholeFame  int    `json:"wholeFame"`
	ClanStatus string `json:"clanStatus"`
	JoinDate   string `json:"joinDate"`
	FkClan     string `json:"fk_clan"`
}

// -------------------- End Structs -------------------- //

// -------------------- Start Main -------------------- //
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Fehler beim Laden der .env-Datei")
	}

	gin.SetMode(gin.ReleaseMode)

	logFile, err := os.OpenFile("gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Fehler beim Öffnen der Log-Datei:", err)
	}
	defer logFile.Close()

	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	router := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Request environment variables
	accessToken = os.Getenv("ACCESS_TOKEN")
	clanTag = os.Getenv("CLAN_TAG")
	encodedClanTag = url.QueryEscape(clanTag)

	router.GET("/clan", getClanHandler)
	router.GET("/members", getMembersHandler)
	router.GET("/currentriverrace", getCurrentRiverRaceHandler)
	router.GET("/riverracelog", getRiverRaceLogHandler)
	router.GET("/person", getPersonHandler)

	log.Printf("Server läuft auf Port %s", port)
	log.Fatal(router.Run(":" + port))
}

// -------------------- End Main -------------------- //

// -------------------- Start Requests -------------------- //

// Get Clan Information
func getClanHandler(c *gin.Context) {
	url := "https://api.clashroyale.com/v1/clans/" + encodedClanTag
	response, err := makeRequest(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}
	defer response.Body.Close()

	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	c.JSON(http.StatusOK, data)
}

// Get Members of Clan
func getMembersHandler(c *gin.Context) {
	url := "https://api.clashroyale.com/v1/clans/" + encodedClanTag + "/members"
	response, err := makeRequest(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}
	defer response.Body.Close()

	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	c.JSON(http.StatusOK, data)
}

// Get Current Riverrace information
func getCurrentRiverRaceHandler(c *gin.Context) {
	url := "https://api.clashroyale.com/v1/clans/" + encodedClanTag + "/currentriverrace"
	response, err := makeRequest(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}
	defer response.Body.Close()

	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func getRiverRaceLogHandler(c *gin.Context) {
	url := "https://api.clashroyale.com/v1/clans/" + encodedClanTag + "/riverracelog"
	response, err := makeRequest(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}
	defer response.Body.Close()

	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func makeRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Access-Control-Allow-Origin", "*")

	client := &http.Client{}
	return client.Do(req)
}

// -------------------- End Requests -------------------- //

// -------------------- Start Database -------------------- //

// Connect to MySQL database
func connectToDatabase() (*sql.DB, error) {
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getPersonHandler(c *gin.Context) {
	db, err := connectToDatabase()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}
	defer db.Close()

	var person Person
	err = db.QueryRow("SELECT * FROM person WHERE tag = '#2Y9VQVJ9'").Scan(&person.Tag, &person.WholeFame, &person.ClanStatus, &person.JoinDate, &person.FkClan)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	personJSON, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ein Fehler ist aufgetreten"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(personJSON))
}

// -------------------- End Database -------------------- //
