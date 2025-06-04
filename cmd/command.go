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

type CommandFactory struct {
}

func (f CommandFactory) CreateCommand(name string) (Command, error) {
	switch name {
	case "summarise":
		return &SummariseCommand{}, nil
	case "analyze":
		return &AnalyseCommand{}, nil
	case "process":
		return &ProcessCommand{}, nil
	default:
		return nil, fmt.Errorf("unknown command name: %s", name)
	}
}
