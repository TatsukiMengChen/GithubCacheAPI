package config

import (
    "github.com/joho/godotenv"
    "log"
)

func LoadEnv() {
	log.Println("Loading environment variables...")
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found or failed to load .env")
    }
}
