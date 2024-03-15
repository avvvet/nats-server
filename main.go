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

	//subscribe to topic
	sub, _ := nc.SubscribeSync("media.*")

	//publish to topic
	nc.Publish("media.video", []byte("video track 001"))

	//consume the message
	msg, _ := sub.NextMsg(time.Millisecond * 10)
	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)

	// request reply

	natsub, _ := nc.Subscribe("greet.*", func(msg *nats.Msg) {
		name := msg.Subject[6:]
		fmt.Println("from request ", string(msg.Data))
		msg.Respond([]byte("hello, " + name))

	})

	rep, _ := nc.Request("greet.awet", []byte("hellow from request"), time.Second)
	fmt.Println(string(rep.Data))

	natsub.Unsubscribe()

}
