package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalln(err)
	}

	defer nc.Close()

	for i := range 200 {
		time.Sleep(500 * time.Millisecond)
		nc.Publish("greeting", []byte(fmt.Sprintf("Message #: %d, Hello World!", i)))
	}

	nc.Flush()
}
