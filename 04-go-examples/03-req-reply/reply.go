package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

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

	var i = 0
	var userName UserName
	nc.Subscribe("greeting", func(msg *nats.Msg) {
		i++
		err := json.Unmarshal(msg.Data, &userName)
		if err != nil {
			log.Fatalln(err)
		}
		msg.Respond([]byte(fmt.Sprintf("Hi %s %s, Have a nice day!", userName.FirstName, userName.LastName)))
	})
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]", "greeting")

	// Setup the interrupt handler to drain so we don't miss
	// requests when scaling down.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println()
	log.Printf("Draining...")
	nc.Drain()
	log.Fatalf("Exiting")

}

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
}
