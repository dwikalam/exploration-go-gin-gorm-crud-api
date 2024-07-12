package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	validateEnvVariables()
}

func validateEnvVariables() {
	validateDbEnvVariables()
	validateServerEnvVariables()
}

func validateDbEnvVariables() {
	if os.Getenv("DB_URL") == "" {
		log.Fatal("Error. DB_URL is missing")
	}
}

func validateServerEnvVariables() {
	if os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Error. SERVER_PORT is missing")
	}
}
