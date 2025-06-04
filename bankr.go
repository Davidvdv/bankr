package main

import (
	"bankr/cmd"
	"bankr/internal"
	"bankr/internal/classification"
	"bankr/internal/io"
	"bankr/internal/model"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const csvFileExtension = ".csv"
const defaultDir = "internal/resources"

type Application interface {
	Go()
}

type BankrApp struct {
	resourcesDir         string
	csvFileReader        io.FileReader
	transactionProcessor internal.Processor
}

func ApplicationFactory() Application {
	app := &BankrApp{
		resourcesDir:         defaultDir,
		csvFileReader:        &io.CsvFileReader{},
		transactionProcessor: &internal.TransactionProcessor{},
	}
	return app
}

func (b *BankrApp) Go() {
	fmt.Print("### Bankr CLI! ###\n\n")

	registry := cmd.NewCommandRegistry()
	registry.Register("summarise", &cmd.SummariseCommand{})
	registry.Register("process", &cmd.ProcessCommand{})

	// Handle command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <command> [args...]")
		registry.ListCommands()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	// Execute the command
	if err := registry.Execute(command, args); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

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
