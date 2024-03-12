package main

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

func main() {

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = "nats://localhost:4222"
	}

	nc, err := nats.Connect(url)
	if err != nil {
		log.Printf("Error: unbale to connect to NATS server %v", err)
	}

	log.Printf("Connected to NATS server at %s", url)

	// Close the NATS connection when done
	defer nc.Close()
}
