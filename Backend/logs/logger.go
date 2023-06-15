package logs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Log message to file
func LogMessage(logType string, message string) {

	// Open log file in append mode
	logFile, err := os.OpenFile(filepath.Join("logs", "logfile.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error while creating log file: " + err.Error())
	}

	// Close log file
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Println("Error while closing log file: " + err.Error())
		}
	}(logFile)

	// Log to file
	logger := log.New(logFile, "", log.Ldate|log.Ltime)
	logString := fmt.Sprintf("[%s] %s", logType, message)
	logger.Println(logString)
}
