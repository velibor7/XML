package config

import "os"

type Config struct {
	Port               string
	AuthenticationHost string
	AuthenticationPort string
	ProfileHost        string
	ProfilePort        string
	ConnectionPort     string
	ConnectionHost     string
	PostHost           string
	PostPort           string
	JobHost            string
	JobPort            string
}

func NewConfig() *Config {
	return &Config{
		Port:               os.Getenv("GATEWAY_PORT"),
		AuthenticationHost: os.Getenv("AUTHENTICATION_SERVICE_HOST"),
		AuthenticationPort: os.Getenv("AUTHENTICATION_SERVICE_PORT"),
		ProfileHost:        os.Getenv("PROFILE_SERVICE_HOST"),
		ProfilePort:        os.Getenv("PROFILE_SERVICE_PORT"),
		PostHost:           os.Getenv("POST_SERVICE_HOST"),
		PostPort:           os.Getenv("POST_SERVICE_PORT"),
		JobHost:            os.Getenv("JOB_SERVICE_HOST"),
		JobPort:            os.Getenv("JOB_SERVICE_PORT"),
		ConnectionHost:     os.Getenv("CONNECTION_SERVICE_HOST"),
		ConnectionPort:     os.Getenv("CONNECTION_SERVICE_PORT"),
	}
}
