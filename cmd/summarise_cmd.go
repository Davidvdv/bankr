package cmd

import (
	"bankr/internal"
	"bankr/internal/io"
	"bankr/internal/model"
	"fmt"
)

type SummariseCommand struct {
	directoryReader io.DirectoryReader
	fileReader      io.FileReader
}

func (s *SummariseCommand) Execute(args []string) error {
	filePaths, err := s.directoryReader.Ls(args[0])
	if err != nil {
		return fmt.Errorf("error listing files: %v", err)
	}
	linesOfFiles := s.fileReader.ReadLinesOfFiles(filePaths)
	transactions := model.BuildTransactions(linesOfFiles)

	printSummary(transactions, len(filePaths))
	return nil
}

func (s *SummariseCommand) Description() string {
	return "Provides a summary of the transactions in the CSV files"
}

func printSummary(transactions []*model.Transaction, numberOfAccounts int) {
	summary := model.BuildSummary(transactions, numberOfAccounts)
	fmt.Printf("Summary:\n%s\n", internal.PrettyJson(summary))
}
