package main

import (
	"bankr/internal/model"
	"encoding/json"
	"fmt"
	"bankr/internal/io"
	"os"
	"path/filepath"
	"strings"
)

const CSV_FILE_EXTENSION = ".csv"
const DEFAULT_DIR = "resources"

type Application interface {
	Go()
}

type BankrCli struct {
	resourcesDir string
	csvFileReader io.FileReader
}

func ApplicationFactory() Application {
	app := &BankrCli{
		resourcesDir: DEFAULT_DIR,
		csvFileReader: &io.CsvFileReader{},
	}
	return app
}

func (a *BankrCli) Go() {
	filepaths, err := allCsvFilesInDir(DEFAULT_DIR)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	printFileDetails(filepaths)

	allEntries := a.csvFileReader.ReadEntriesOfFiles(filepaths)

	transactions := model.BuildTransactions(allEntries)
	summary := model.BuildSummary(transactions)
	jsonSummary, _ := json.MarshalIndent(summary, "", "  ")
	fmt.Printf("Summary:\n%s\n", string(jsonSummary))
}

func printFileDetails(filepaths []string) {
	fmt.Printf("Found %d CSV files\n", len(filepaths))
	for _, filePath := range filepaths {
		fmt.Println(filePath)
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
		if !strings.HasSuffix(file.Name(), CSV_FILE_EXTENSION) {
			continue
		}
		filePath := filepath.Join(dirPath, file.Name())
		filePaths = append(filePaths, filePath)
	}

	return filePaths, nil
}
