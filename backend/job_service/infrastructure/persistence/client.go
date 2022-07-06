package persistence

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetClient(uri, username, password string) (*neo4j.Driver, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth("neo4j", "password", ""))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// mozda ne treba close da se uradi
	//defer driver.Close()
	return &driver, nil
}
