package config

import "os"

type Config struct {
	Port                        string
	AuthDBHost                  string
	AuthDBPort                  string
	ProfileHost                 string
	ProfilePort                 string
	NatsHost                    string
	NatsPort                    string
	NatsUser                    string
	NatsPass                    string
	CreateProfileCommandSubject string
	CreateProfileReplySubject   string
	UpdateProfileCommandSubject string
	UpdateProfileReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                        os.Getenv("AUTH_SERVICE_PORT"),
		AuthDBHost:                  os.Getenv("AUTH_DB_HOST"),
		AuthDBPort:                  os.Getenv("AUTH_DB_PORT"),
		ProfileHost:                 os.Getenv("PROFILE_SERVICE_HOST"),
		ProfilePort:                 os.Getenv("PROFILE_SERVICE_PORT"),
		NatsHost:                    os.Getenv("NATS_HOST"),
		NatsPort:                    os.Getenv("NATS_PORT"),
		NatsUser:                    os.Getenv("NATS_USER"),
		NatsPass:                    os.Getenv("NATS_PASS"),
		CreateProfileCommandSubject: os.Getenv("CREATE_PROFILE_COMMAND_SUBJECT"),
		CreateProfileReplySubject:   os.Getenv("CREATE_PROFILE_REPLY_SUBJECT"),
		UpdateProfileCommandSubject: os.Getenv("UPDATE_PROFILE_COMMAND_SUBJECT"),
		UpdateProfileReplySubject:   os.Getenv("UPDATE_PROFILE_REPLY_SUBJECT"),
	}
}
