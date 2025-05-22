package main

import (
	"bankr/model"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const CSV_FILE_EXTENSION = ".csv"
const DEFAULT_DIR = "csv"

type Application interface {
	Go()
}

type BankrCli struct {
}

func ApplicationFactory() Application {
	app := &BankrCli{}
	return app
}

func (a *BankrCli) Go() {
	filepaths, err := allCsvFilesInDir(DEFAULT_DIR)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Found %d CSV files\n", len(filepaths))
	for _, filePath := range filepaths {
		fmt.Println(filePath)
	}

	readAllCsvFiles(filepaths)
}

func readAllCsvFiles(filepaths []string) {
	ch := make(chan [][]string, len(filepaths))

	for _, filePath := range filepaths {
		go readFile(filePath, ch)
	}

	var allEntries [][]string
	for range filepaths {
		entries := <-ch
		allEntries = append(allEntries, entries...)
	}

	transactions := model.BuildTransactions(allEntries)
	summary := model.BuildSummary(transactions)
	jsonSummary, _ := json.MarshalIndent(summary, "", "  ")
	fmt.Printf("Summary:\n%s\n", string(jsonSummary))
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

func readFile(filePath string, ch chan [][]string) {
	fmt.Printf("Reading file %s\n", filePath)
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("could not open file %s: %v\n", filePath, err)
		ch <- make([][]string, 0)
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	_, _ = reader.Read() // Skip the header row

	var lines [][]string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("could not read CSV: %v", err)
			ch <- lines
			return
		}
		lines = append(lines, line)
	}

	fmt.Printf("Successfully read %d lines from CSV\n", len(lines))
	ch <- lines
}
