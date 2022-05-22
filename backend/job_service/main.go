package main

import (
	"github.com/velibor7/XML/job_service/startup"
	"github.com/velibor7/XML/job_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
