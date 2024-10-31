package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("Process Id :", os.Getpid())
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalln(err)
	}

	nc.Subscribe("greeting", func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
	})

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt)

	<-shutdownCh
	fmt.Println("Client shutdown!")
}
