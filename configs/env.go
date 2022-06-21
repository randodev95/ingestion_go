package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func FetchEnv(requirement string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading Environment Variables")
	}

	return os.Getenv(requirement)
}
