package cmd

import (
	"bankr/internal"
	"bankr/internal/io"
	"bankr/internal/model"
	"fmt"
	"regexp"
	"strings"
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
		return strings.TrimSpace(sanitiseTransactionDetails(t.Details) + " " + strings.TrimSuffix(t.Code, " C"))
	})
	internal.PrettyPrintJson(descriptions)
	//internal.PrettyPrintJson(classification.AnalyzeDescriptions(descriptions))
	return nil
}

func (a *AnalyseCommand) Description() string {
	return "Analyses the type, description and code of the transactions in the CSV files"
}

func sanitiseTransactionDetails(details string) string {
	return regexp.MustCompile("4835-.*-2708").ReplaceAllString(strings.TrimSuffix(details, " Df"), "")
}
