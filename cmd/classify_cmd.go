package cmd

import (
	"bankr/internal/classification"
	"bankr/internal/io"
	"bankr/internal/model"
	"fmt"
)

type ClassifyCommand struct {
	directoryReader io.DirectoryReader
	fileReader      io.FileReader
	classifier      classification.Classifier
}

func (c *ClassifyCommand) Execute(args []string) error {
	fmt.Println("=> Classify")
	filePaths, err := c.directoryReader.Ls(args[0])
	if err != nil {
		return fmt.Errorf("error listing files: %v", err)
	}
	linesOfFiles := c.fileReader.ReadLinesOfFiles(filePaths)
	transactions := model.BuildTransactions(linesOfFiles)

	descriptions := model.Map(transactions, func(t *model.Transaction) string {
		return t.Details + t.Code
	})

	c.classifier.Classify(descriptions)

	return nil
}

func (c *ClassifyCommand) Description() string {
	return "Classifies the transactions in the CSV files"
}
