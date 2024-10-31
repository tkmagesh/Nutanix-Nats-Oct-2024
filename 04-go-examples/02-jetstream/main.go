package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {

	/* Connect to NATS */
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)
	defer cancel()

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	/* Stream creation */
	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}
	s, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "TEST_STREAM",
		Subjects: []string{"FOO.*"},
	})
	if err != nil {
		log.Fatal(err)
	}

	/* publish messages to the stream */
	go endlessPublish(ctx, nc, js)

	/* Consume the messages published to the stream */
	cons, err := s.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "TestConsumerMessages",
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Fatal(err)
	}

	it, err := cons.Messages(jetstream.PullMaxMessages(1))
	if err != nil {
		log.Fatal(err)
	}
	for {
		msg, err := it.Next()
		if err != nil {
			fmt.Println("next err: ", err)
		}
		fmt.Println(string(msg.Data()))
		msg.Ack()
	}
}

func endlessPublish(ctx context.Context, nc *nats.Conn, js jetstream.JetStream) {
	var i int
	for {
		time.Sleep(500 * time.Millisecond)
		if nc.Status() != nats.CONNECTED {
			continue
		}
		if _, err := js.Publish(ctx, "FOO.TEST1", []byte(fmt.Sprintf("msg %d", i))); err != nil {
			fmt.Println("pub error: ", err)
		}
		i++
	}
}
