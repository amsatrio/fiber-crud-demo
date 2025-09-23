package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("error: %v\n", err)
		log.Fatal("Error loading .env file")
		return
	}
}
