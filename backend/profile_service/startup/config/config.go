package config

import "os"

type Config struct {
	Port                        string
	ProfileDBHost               string
	ProfileDBPort               string
	CommentHost                 string
	CommentPort                 string
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
		Port:                        os.Getenv("PROFILE_SERVICE_PORT"),
		ProfileDBHost:               os.Getenv("PROFILE_DB_HOST"),
		ProfileDBPort:               os.Getenv("PROFILE_DB_PORT"),
		NatsHost:                    os.Getenv("NATS_HOST"),
		NatsPort:                    os.Getenv("NATS_PORT"),
		NatsUser:                    os.Getenv("NATS_USER"),
		NatsPass:                    os.Getenv("NATS_PASS"),
		CommentHost:                 os.Getenv("COMMENT_SERVICE_HOST"),
		CommentPort:                 os.Getenv("COMMENT_SERVICE_PORT"),
		CreateProfileCommandSubject: os.Getenv("CREATE_PROFILE_COMMAND_SUBJECT"),
		CreateProfileReplySubject:   os.Getenv("CREATE_PROFILE_REPLY_SUBJECT"),
		UpdateProfileCommandSubject: os.Getenv("UPDATE_PROFILE_COMMAND_SUBJECT"),
		UpdateProfileReplySubject:   os.Getenv("UPDATE_PROFILE_REPLY_SUBJECT"),
	}
}
