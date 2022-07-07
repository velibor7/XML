package config

import "os"

type Config struct {
	Port          string
	JobDBHost     string
	JobDBPort     string
	Neo4jUri      string
	Neo4jUsername string
	Neo4jPassword string
	NatsHost      string
	NatsPort      string
	NatsUser      string
	NatsPass      string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("JOB_SERVICE_PORT"),
		Neo4jUri:      os.Getenv("NEO4J_URI"),
		Neo4jUsername: os.Getenv("NEO4J_USERNAME"),
		Neo4jPassword: os.Getenv("NEO4J_PASSWORD"),
		NatsHost:      os.Getenv("NATS_HOST"),
		NatsPort:      os.Getenv("NATS_PORT"),
		NatsUser:      os.Getenv("NATS_USER"),
		NatsPass:      os.Getenv("NATS_PASS"),
	}
}
