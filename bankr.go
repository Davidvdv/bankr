package main

import (
	"bankr/internal"
	"bankr/internal/io"
	"bankr/internal/model"
	"fmt"
	"github.com/fatih/color"
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
	c := color.New(color.FgYellow, color.Bold)
	_, _ = c.Print("### Bankr CLI! ###\n\n")

	filePaths, err := allCsvFilesInDir(defaultDir)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	printFileDetails(filePaths)

	allEntries := b.csvFileReader.ReadEntriesOfFiles(filePaths)
	transactions := model.BuildTransactions(allEntries)

	printSummary(transactions, len(filePaths))

	b.transactionProcessor.Process(transactions)
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
