package main

import (
	"bankr/cmd"
	"fmt"
	"os"
)

type Application interface {
	Go()
}

type BankrApp struct {
}

func ApplicationFactory() Application {
	return &BankrApp{}
}

func (b *BankrApp) Go() {
	fmt.Print("### Bankr CLI! ###\n\n")

	summariseCmd, err := cmd.CreateCommand("summarise")
	processCmd, err := cmd.CreateCommand("process")
	analyseCmd, err := cmd.CreateCommand("analyse")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	registry := cmd.NewCommandRegistry()
	registry.Register("summarise", summariseCmd)
	registry.Register("process", processCmd)
	registry.Register("analyse", analyseCmd)

	if len(os.Args) < 2 {
		fmt.Println("Usage: program <command> [args...]")
		registry.ListCommands()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	if err := registry.Execute(command, args); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
