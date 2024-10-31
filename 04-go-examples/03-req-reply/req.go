package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type UserName struct {
	FirstName string
	LastName  string
}

func main() {

	/* Connect to NATS */

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	payloadData := UserName{
		FirstName: "Magesh",
		LastName:  "Kuppan",
	}

	payload, err := json.Marshal(payloadData)
	if err != nil {
		log.Fatalln(err)
	}

	msg, err := nc.Request("greeting", payload, 5*time.Second)
	if err != nil {
		if nc.LastError() != nil {
			log.Fatalf("%v for request", nc.LastError())
		}
		log.Fatalf("%v for request", err)
	}

	log.Printf("Published [%s] : '%s'", "greeting", payload)
	log.Printf("Received  [%v] : '%s'", msg.Subject, string(msg.Data))
}
