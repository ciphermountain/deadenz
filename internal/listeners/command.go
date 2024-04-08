package listeners

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	deadenz "github.com/ciphermountain/deadenz/pkg"
)

type CommandEvent struct {
	reader     *bufio.Reader
	commands   map[string]deadenz.CommandType
	chCommands chan deadenz.CommandType
	chPrompt   chan struct{}

	mu             sync.Mutex
	defaultCommand deadenz.CommandType
}

func NewCommandEvent(
	defaultCommand deadenz.CommandType,
	commands map[string]deadenz.CommandType,
) *CommandEvent {
	listener := &CommandEvent{
		reader:         bufio.NewReader(os.Stdin),
		chCommands:     make(chan deadenz.CommandType, 1),
		chPrompt:       make(chan struct{}, 1),
		defaultCommand: defaultCommand,
	}

	go listener.run()

	return listener
}

func (e *CommandEvent) Next() <-chan deadenz.CommandType {
	e.chPrompt <- struct{}{}

	return e.chCommands
}

func (e *CommandEvent) SetDefaultCommand(command deadenz.CommandType) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.defaultCommand = command
}

func (e *CommandEvent) run() {
	for {
		<-e.chPrompt

		e.mu.Lock()
		def := e.defaultCommand
		e.mu.Unlock()

		var defStr string
		for key, value := range e.commands {
			if value == def {
				defStr = string(key)

				break
			}
		}

		fmt.Printf("Enter command (%s): ", defStr)

		input, err := e.reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)

			continue
		}

		// remove the delimeter from the string
		input = strings.TrimSuffix(input, "\n")

		// set default input
		if len(input) == 0 {
			e.chCommands <- def

			return
		}

		cmd, ok := e.commands[input]
		if !ok {
			fmt.Println("unrecognized command")

			return
		}

		e.chCommands <- cmd
	}
}

var EnglishCommands map[string]deadenz.CommandType = map[string]deadenz.CommandType{
	"spawnin":  deadenz.SpawninCommandType,
	"walk":     deadenz.WalkCommandType,
	"backpack": deadenz.BackpackCommandType,
	"xp":       deadenz.XPCommandType,
	"currency": deadenz.CurrencyCommandType,
	"exit":     deadenz.ExitCommandType,
	"quit":     deadenz.ExitCommandType,
}
