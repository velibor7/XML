package main

import (
	"fmt"

	startup "github.com/velibor7/XML/connection_service/startup"
	cfg "github.com/velibor7/XML/connection_service/startup/config"
)

func main() {
	fmt.Println("Hello world")

	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
