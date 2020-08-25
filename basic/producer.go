package main

import (
	"context"
	"fmt"
	log "github.com/apache/pulsar/pulsar-client-go/logutil"
	"github.com/apache/pulsar/pulsar-client-go/pulsar"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:       "pulsar://localhost:6650",
		IOThreads: 5,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})
	if err != nil {
		log.Fatal(err)
	}

	defer producer.Close()

	ctx := context.Background()

	for i := 0; i < 10; i++ {
		if err := producer.Send(ctx, pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprintf("hello-%d", i)),
		}); err != nil {
			log.Fatal(err)
		}
	}
}
