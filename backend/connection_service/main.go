package main

import (
	"github.com/velibor7/XML/connection_service/startup"
	cfg "github.com/velibor7/XML/connection_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
