package main

import (
	"github.com/velibor7/XML/authentication_service/startup"
	"github.com/velibor7/XML/authentication_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
