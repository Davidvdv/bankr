package cmd

import "fmt"

type Command interface {
	Execute(args []string)
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
	cmd.Execute(args)
	return nil
}

func (r *CommandRegistry) ListCommands() {
	fmt.Println("Available commands:")
	for name, cmd := range r.commands {
		fmt.Printf("  %s - %s\n", name, cmd.Description())
	}
}

type SummariseCommand struct {
}

func (s SummariseCommand) Execute([]string) {
	//TODO implement me
	panic("implement me")
}

func (s SummariseCommand) Description() string {
	return "Provides a summary of the transactions in the CSV files"
}

type AnalyzeCommand struct {
}

func (a AnalyzeCommand) Execute([]string) {
	//TODO implement me
	panic("implement me")
}

func (a AnalyzeCommand) Description() string {
	return "Analyzes the type, description and code of the transactions in the CSV files"
}

type ProcessCommand struct{}

func (p ProcessCommand) Execute([]string) {
	//TODO implement me
	panic("implement me")
}

func (p ProcessCommand) Description() string {
	return "Process the transactions in the CSV files and generate a summary of the transactions and their categories."
}
