package cmd

import (
	"bankr/internal"
	"bankr/internal/io"
	"fmt"
)

type Command interface {
	Execute(args []string) error
	Description() string
}

type CommandRegistry struct {
	commands map[string]Command
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		commands: make(map[string]Command),
	}
}

func (r *CommandRegistry) Register(name string, command Command) {
	r.commands[name] = command
}

func (r *CommandRegistry) Execute(name string, args []string) error {
	cmd, exists := r.commands[name]
	if !exists {
		return fmt.Errorf("command '%s' not found", name)
	}
	err := cmd.Execute(args)
	if err != nil {
		return err
	}
	return nil
}

func (r *CommandRegistry) ListCommands() {
	fmt.Println("Available commands:")
	for name, cmd := range r.commands {
		fmt.Printf("  %s - %s\n", name, cmd.Description())
	}
}

func CreateCommand(name string) (Command, error) {
	switch name {
	case "summarise":
		return &SummariseCommand{
			directoryReader: &io.LocalDirectoryReader{},
			fileReader:      &io.CsvFileReader{},
		}, nil
	case "analyse":
		return &AnalyseCommand{
			directoryReader: &io.LocalDirectoryReader{},
			fileReader:      &io.CsvFileReader{},
		}, nil
	case "process":
		return &ProcessCommand{
			directoryReader:      &io.LocalDirectoryReader{},
			fileReader:           &io.CsvFileReader{},
			transactionProcessor: &internal.TransactionProcessor{},
		}, nil
	case "classify":
		return &ClassifyCommand{
			directoryReader: &io.LocalDirectoryReader{},
			fileReader:      &io.CsvFileReader{},
		}, nil
	default:
		return nil, fmt.Errorf("unknown command name: %s", name)
	}
}
