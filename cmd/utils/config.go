package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Failed to load environment variables")
	}
	return os.Getenv(key)
}
