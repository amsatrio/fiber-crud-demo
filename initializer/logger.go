package initializer

import (
	"fiber-crud-demo/util"
	"fmt"
	"os"

	"log"
	"time"
)

func LoggerInit() {

	// ASCII
	asciiArt := `
		____
	   < hi there >
		----
			 \   ^__^
			  \  (oo)\_______
				 (__)\       )\/\
					 ||----w |
					 ||     ||
		`

	logType := os.Getenv("LOG_TYPE")
	if logType == "stdout" {
		log.Println(asciiArt)
		return
	}

	logDir := os.Getenv("LOG_DIR")
	if logDir == "" {
		logDir = "logs"
	}

	// clear log
	if os.Getenv("LOG_CLEAR") == "true" {
		err := util.RemoveAll(logDir)
		if err != nil {
			fmt.Println("delete log error: " + err.Error())
		}
	}

	requestTime := time.Now().Format("2006-01-02")

	// Create a log directory if it doesn't exist
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, os.ModePerm)
		if err != nil {
			log.Fatal("Failed to create log directory")
			return
		}
	}

	// Create a log file if it doesn't exist
	logFileName := logDir + "/app_" + requestTime + ".log"
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		defer func(logFile *os.File) {
			err := logFile.Close()
			if err != nil {
				log.Println(err)
			}
		}(logFile)
		log.Fatal("Failed to create or open log file:", err)
		return
	}

	// Set the log file as the output for the standard logger
	log.SetOutput(logFile)

	log.Println(asciiArt)
}
