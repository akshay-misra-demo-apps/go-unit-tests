package main

import (
	"fmt"

	"github.com/misraak/app/cmd"
)

// main function of this microservice worker.
// The job of this worker it to listen to kafka messages, and add records to/update records in/delete records from InMemory DB.
// This worker support the following databases: HazelCast.
func main() {

	err := cmd.Init()
	if err != nil {
		// Print the error occured while starting the server.
		fmt.Printf("%+v\n", err.Error())
	}
	return
}
