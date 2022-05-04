package config

import "os"

type Config struct {
	Port           string
	SecurityDBHost string
	SecurityDBPort string
	NatsHost       string
	NatsPort       string
	NatsUser       string
	NatsPass       string
}

func NewConfig() *Config {
	return &Config{
		Port:           "8001",
		SecurityDBHost: "security_db",
		SecurityDBPort: "27017",
		NatsHost:       os.Getenv("NATS_HOST"),
		NatsPort:       os.Getenv("NATS_PORT"),
		NatsUser:       os.Getenv("NATS_USER"),
		NatsPass:       os.Getenv("NATS_PASS"),
	}
}
