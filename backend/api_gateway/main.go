package main

import (
	"github.com/velibor7/XML/api_gateway/startup"
	"github.com/velibor7/XML/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
