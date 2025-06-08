package cmd

import (
	"bankr/internal"
	"bankr/internal/classification"
	"bankr/internal/io"
	"bankr/internal/model"
	"fmt"
)

type AnalyseCommand struct {
	directoryReader io.DirectoryReader
	fileReader      io.FileReader
}

func (a *AnalyseCommand) Execute(args []string) error {
	filePaths, err := a.directoryReader.Ls(args[0])
	if err != nil {
		return fmt.Errorf("error listing files: %v", err)
	}
	linesOfFiles := a.fileReader.ReadLinesOfFiles(filePaths)
	transactions := model.BuildTransactions(linesOfFiles)

	descriptions := model.Map(transactions, func(t *model.Transaction) string {
		return t.Details + t.Code
	})
	internal.PrettyPrintJson(descriptions)
	internal.PrettyPrintJson(classification.AnalyzeDescriptions(descriptions))
	return nil
}

func (a *AnalyseCommand) Description() string {
	return "Analyses the type, description and code of the transactions in the CSV files"
}
