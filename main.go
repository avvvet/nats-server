package main

import (
	"log"
	"os"
)

func main() {

	url := os.Getenv("NATS_URL")
	log.Println(url)
}
