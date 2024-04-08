package listeners

import (
	"context"
	"log"

	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/parse"
	"github.com/ciphermountain/deadenz/pkg/service/multiverse"
)

type MultiverseMessageListener struct {
	client *multiverse.MultiverseClient

	reader   *multiverse.EventsReader
	chEvents chan components.Event
}

func NewMultiverseMessageListener(client *multiverse.MultiverseClient) (*MultiverseMessageListener, error) {
	reader, err := client.NewEventsStreamReader(context.Background())
	if err != nil {
		return nil, err
	}

	listener := &MultiverseMessageListener{
		client:   client,
		reader:   reader,
		chEvents: make(chan components.Event, 100),
	}

	go listener.run()

	return listener, nil
}

func (e *MultiverseMessageListener) Next() <-chan components.Event {
	return e.chEvents
}

func (e *MultiverseMessageListener) Close() error {
	return e.reader.Close()
}

func (e *MultiverseMessageListener) run() {
	for {
		event, err := e.reader.Next()
		if err != nil {
			log.Printf("error reading message from grpc: %s", err.Error())

			break
		}

		evt, err := parse.DecodeJSONEvent(event.Data)
		if err != nil {
			log.Printf("error reading message from grpc: %s", err.Error())

			continue
		}
		log.Println("event detected from metaverse")

		e.chEvents <- evt
	}
}
