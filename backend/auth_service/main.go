package main

import (
	"github.com/velibor7/XML/auth_service/startup"
	"github.com/velibor7/XML/auth_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
