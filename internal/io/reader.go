package io

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const csvFileExtension = ".csv"

type DirectoryReader interface {
	Ls(dirPath string) ([]string, error)
}

type LocalDirectoryReader struct {
}

func (l *LocalDirectoryReader) Ls(dirPath string) ([]string, error) {
	absPath, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, fmt.Errorf("error getting absolute path: %v", err)
	}
	files, err := os.ReadDir(absPath)
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

	printFileDetails(filePaths)

	return filePaths, nil
}

type FileReader interface {
	ReadLinesOfFiles(filePaths []string) [][]string
}

type CsvFileReader struct{}

func (c *CsvFileReader) ReadLinesOfFiles(filePaths []string) [][]string {
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
	fmt.Printf("=> Reading file %s\n", filePath)
	csvFile, err := os.Open(filePath)
	defer func(csvFile *os.File) {
		_ = csvFile.Close()
	}(csvFile)

	if err != nil {
		fmt.Printf("could not open file %s: %v\n", filePath, err)
		ch <- make([][]string, 0)
		return
	}

	reader := csv.NewReader(csvFile)
	_, _ = reader.Read() // Skip the header row

	var lines = make([][]string, 0)
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

	fmt.Printf("=> Successfully read %d lines from CSV\n", len(lines))

	ch <- lines
}

func printFileDetails(filePaths []string) {
	fmt.Printf("=> Found %d CSV files\n", len(filePaths))
	for _, filePath := range filePaths {
		fmt.Println(" *", filePath)
	}
}
