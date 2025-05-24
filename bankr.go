package main

import (
	"bankr/internal"
	"bankr/internal/io"
	"bankr/internal/model"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/fatih/color"
)

const CSV_FILE_EXTENSION = ".csv"
const DEFAULT_DIR = "internal/resources"

type Application interface {
	Go()
}

type BankrCli struct {
	resourcesDir         string
	csvFileReader        io.FileReader
	transactionProcessor internal.Processor
}

func ApplicationFactory() Application {
	app := &BankrCli{
		resourcesDir:         DEFAULT_DIR,
		csvFileReader:        &io.CsvFileReader{},
		transactionProcessor: &internal.TransactionProcessor{},
	}
	return app
}
func (a *BankrCli) Go() {
	c := color.New(color.FgYellow, color.Bold)
	c.Print("### Bankr CLI! ###\n\n")
	
	filepaths, err := allCsvFilesInDir(DEFAULT_DIR)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	printFileDetails(filepaths)

	allEntries := a.csvFileReader.ReadEntriesOfFiles(filepaths)

	transactions := model.BuildTransactions(allEntries)
	printSummary(transactions, len(filepaths))

	a.transactionProcessor.Process(transactions)
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
		if !strings.HasSuffix(file.Name(), CSV_FILE_EXTENSION) {
			continue
		}
		filePath := filepath.Join(dirPath, file.Name())
		filePaths = append(filePaths, filePath)
	}

	return filePaths, nil
}

func printFileDetails(filepaths []string) {
	fmt.Printf("=> Found %d CSV files\n", len(filepaths))
	for _, filePath := range filepaths {
		fmt.Println(filePath)
	}
}

func printSummary(transactions []*model.Transaction, numberOfAccounts int) {
	summary := model.BuildSummary(transactions, numberOfAccounts)
	fmt.Printf("Summary:\n%s\n", internal.PrettyJson(summary))
}