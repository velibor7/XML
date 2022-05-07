package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	PostDBHost string
	PostDBPort string
}

func NewConfig() *Config {
	err := CreateEnv()
	if err != nil {
		return nil
	}
	return &Config{
		Port:       os.Getenv("POST_SERVICE_PORT"),
		PostDBHost: os.Getenv("POST_DB_HOST"),
		PostDBPort: os.Getenv("POST_DB_PORT"),
	}
}

func CreateEnv() error {
	if os.Getenv("OS_ENV") != "docker" {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	return nil
}
