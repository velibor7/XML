package main

import (
	"fmt"

	"github.com/velibor7/XML/job_service/startup"
	"github.com/velibor7/XML/job_service/startup/config"
)

func main() {
	config := config.NewConfig()
	fmt.Print("config part\n")
	server := startup.NewServer(config)
	fmt.Print("Server part\n")
	server.Start()
	fmt.Print("Starting ...")
}
