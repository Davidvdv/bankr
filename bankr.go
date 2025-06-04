package main

import (
	"bankr/cmd"
	"bankr/internal"
	"bankr/internal/classification"
	"bankr/internal/model"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const defaultDir = "internal/resources"

type Application interface {
	Go()
}

type BankrApp struct {
	resourcesDir   string
	commandFactory *cmd.CommandFactory
}

func ApplicationFactory() Application {
	app := &BankrApp{
		resourcesDir:   defaultDir,
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

	filePaths, err := allCsvFilesInDir(defaultDir)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	printFileDetails(filePaths)

	allEntries := b.csvFileReader.ReadEntriesOfFiles(filePaths)
	transactions := model.BuildTransactions(allEntries)

	printSummary(transactions, len(filePaths))
	descriptions := model.Map(transactions, func(t *model.Transaction) string {
		return t.Details + t.Code
	})
	internal.PrettyPrintJson(descriptions)
	internal.PrettyPrintJson(classification.AnalyzeDescriptions(descriptions))
	//b.transactionProcessor.Process(transactions)

	// Execute the command
	if err := registry.Execute(command, args); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func allCsvFilesInDir(dirPath string) ([]string, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %v", err)
	}

	filePaths := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), csvFileExtension) {
			continue
		}
		filePath := filepath.Join(dirPath, file.Name())
		filePaths = append(filePaths, filePath)
	}

	return filePaths, nil
}

func printFileDetails(filePaths []string) {
	fmt.Printf("=> Found %d CSV files\n", len(filePaths))
	for _, filePath := range filePaths {
		fmt.Println(filePath)
	}
}

func printSummary(transactions []*model.Transaction, numberOfAccounts int) {
	summary := model.BuildSummary(transactions, numberOfAccounts)
	fmt.Printf("Summary:\n%s\n", internal.PrettyJson(summary))
}
