package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/ciphermountain/deadenz/pkg/events"
	"github.com/ciphermountain/deadenz/pkg/multiverse"
)

type CommandEventListener struct {
	reader     *bufio.Reader
	chCommands chan string
	chPrompt   chan struct{}

	mu             sync.Mutex
	defaultCommand string
}

func NewCommandEventListener(defaultCommand string) *CommandEventListener {
	listener := &CommandEventListener{
		reader:         bufio.NewReader(os.Stdin),
		chCommands:     make(chan string, 1),
		chPrompt:       make(chan struct{}, 1),
		defaultCommand: defaultCommand,
	}

	go listener.run()

	return listener
}

func (e *CommandEventListener) Next() <-chan string {
	e.chPrompt <- struct{}{}

	return e.chCommands
}

func (e *CommandEventListener) SetDefaultCommand(command string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.defaultCommand = command
}

func (e *CommandEventListener) run() {
	for {
		<-e.chPrompt

		e.mu.Lock()
		def := e.defaultCommand
		e.mu.Unlock()

		fmt.Printf("Enter command (%s): ", def)

		input, err := e.reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)

			continue
		}

		// remove the delimeter from the string
		input = strings.TrimSuffix(input, "\n")

		// set default input
		if len(input) == 0 {
			input = def
		}

		e.chCommands <- input
	}
}

type MultiverseMessageListener struct {
	client *multiverse.MultiverseClient

	reader   *multiverse.EventsReader
	chEvents chan events.Event
}

func NewMultiverseMessageListener(client *multiverse.MultiverseClient) (*MultiverseMessageListener, error) {
	reader, err := client.NewEventsStreamReader(context.Background())
	if err != nil {
		return nil, err
	}

	listener := &MultiverseMessageListener{
		client:   client,
		reader:   reader,
		chEvents: make(chan events.Event, 100),
	}

	go listener.run()

	return listener, nil
}

func (e *MultiverseMessageListener) Next() <-chan events.Event {
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

		evt, err := events.DecodeJSONEvent(event.Data)
		if err != nil {
			log.Printf("error reading message from grpc: %s", err.Error())

			continue
		}
		log.Println("event detected from metaverse")

		e.chEvents <- evt
	}
}
