package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	sub, _ := nc.SubscribeSync("media.*")

	nc.Publish("media.video", []byte("video track 001"))

	msg, _ := sub.NextMsg(time.Millisecond * 10)
	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)
}
