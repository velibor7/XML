package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	ProfileDBHost string
	ProfileDBPort string
}

func NewConfig() *Config {
	err := CreateEnv()
	if err != nil {
		return nil
	}
	return &Config{
		Port:          os.Getenv("PROFILE_SERVICE_PORT"),
		ProfileDBHost: os.Getenv("PROFILE_DB_HOST"),
		ProfileDBPort: os.Getenv("PROFILE_DB_PORT"),
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
