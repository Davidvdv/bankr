package cmd

import (
	"bankr/internal"
	"bankr/internal/classification"
	"bankr/internal/io"
	"bankr/internal/model"
	"fmt"
)

type ClassifyCommand struct {
	directoryReader io.DirectoryReader
	fileReader      io.FileReader
}

func (c *ClassifyCommand) Execute(args []string) error {
	filePaths, err := c.directoryReader.Ls(args[0])
	if err != nil {
		return fmt.Errorf("error listing files: %v", err)
	}
	linesOfFiles := c.fileReader.ReadLinesOfFiles(filePaths)
	transactions := model.BuildTransactions(linesOfFiles)

	descriptions := model.Map(transactions, func(t *model.Transaction) string {
		return t.Details + t.Code
	})

	categorizer := classification.NewCategorizer()

	// Add custom rule if needed
	categorizer.AddCustomRule(classification.CategoryFood, `(?i)(my_local_restaurant)`)

	// Categorize
	classifiedTransactions := categorizer.ClassifyTransactions(descriptions, nil)

	internal.PrettyPrintJson(classifiedTransactions)

	// Find transactions that need manual review
	needReview := classification.FindLowConfidenceTransactions(classifiedTransactions, 0.5)

	internal.PrettyPrintJson(needReview)

	return nil
}

func (c *ClassifyCommand) Description() string {
	return "Classifies the transactions in the CSV files"
}
