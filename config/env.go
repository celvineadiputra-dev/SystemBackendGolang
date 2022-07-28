package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading file .env")
	}

	return os.Getenv(key)
}
