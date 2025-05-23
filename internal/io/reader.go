package io

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type FileReader interface {
	ReadEntriesOfFiles(filePaths []string) [][]string
}

type CsvFileReader struct{}

func (c *CsvFileReader) ReadEntriesOfFiles(filePaths []string) [][]string {
	ch := make(chan [][]string, len(filePaths))

	for _, filePath := range filePaths {
		go readFile(filePath, ch)
	}

	var allEntries [][]string
	for range filePaths {
		entries := <-ch
		allEntries = append(allEntries, entries...)
	}

	return allEntries
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
