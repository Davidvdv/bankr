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
	commandFactory *cmd.CommandFactory
}

func ApplicationFactory() Application {
	app := &BankrApp{
		commandFactory: &cmd.CommandFactory{},
	}
	return app
}

func (b *BankrApp) Go() {
	fmt.Print("### Bankr CLI! ###\n\n")

	summariseCmd, err := b.commandFactory.CreateCommand("summarise")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	processCmd, err := b.commandFactory.CreateCommand("process")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	registry := cmd.NewCommandRegistry()
	registry.Register("summarise", summariseCmd)
	registry.Register("process", processCmd)

	if len(os.Args) < 2 {
		fmt.Println("Usage: program <command> [args...]")
		registry.ListCommands()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	//filePaths, err := allCsvFilesInDir(defaultDir)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//}
	//
	//printFileDetails(filePaths)
	//
	//allEntries := b.csvFileReader.ReadEntriesOfFiles(filePaths)
	//transactions := model.BuildTransactions(allEntries)
	//
	//printSummary(transactions, len(filePaths))
	//
	//descriptions := model.Map(transactions, func(t *model.Transaction) string {
	//	return t.Details + t.Code
	//})
	//internal.PrettyPrintJson(descriptions)
	//internal.PrettyPrintJson(classification.AnalyzeDescriptions(descriptions))
	//b.transactionProcessor.Process(transactions)

	if err := registry.Execute(command, args); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
