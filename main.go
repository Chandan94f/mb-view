package main

import (
	"log"

	prepare "github.com/mb-view/cmd"
)

func main() {
	prepare.Prepare()

	if err := prepare.Prepare(); err != nil {
		log.Fatalf("Error preparing the application: %v", err)
	}
	log.Println("Application prepared successfully.")
}
