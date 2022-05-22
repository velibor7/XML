package main

import (
	"fmt"

	"github.com/velibor7/XML/authentication_service/startup"
	cfg "github.com/velibor7/XML/authentication_service/startup/config"
)

func main() {
	fmt.Println("Hello world from auth")

	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()

}
