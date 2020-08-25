package main

import (
	"context"
	"fmt"
	log "github.com/apache/pulsar/pulsar-client-go/logutil"
	"github.com/apache/pulsar/pulsar-client-go/pulsar"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{URL: "pulsar://localhost:6650"})
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "my-topic",
		SubscriptionName: "my-subscription",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()

	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received message  msgId: %s -- content: '%s'\n",
			msg.ID(), string(msg.Payload()))

		consumer.Ack(msg)
	}

}
