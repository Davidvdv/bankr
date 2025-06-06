package cmd

import (
	"bankr/internal"
	"bankr/internal/io"
	"bankr/internal/model"
	"fmt"
)

type ProcessCommand struct {
	directoryReader      io.DirectoryReader
	fileReader           io.FileReader
	transactionProcessor internal.Processor
}

func (p *ProcessCommand) Execute(args []string) error {
	filePaths, err := p.directoryReader.Ls(args[0])
	if err != nil {
		return fmt.Errorf("error listing files: %v", err)
	}
	linesOfFiles := p.fileReader.ReadLinesOfFiles(filePaths)
	transactions := model.BuildTransactions(linesOfFiles)
	p.transactionProcessor.Process(transactions)
	return nil
}

func (p *ProcessCommand) Description() string {
	return "Process the transactions in the CSV files and generate a summary of the transactions and their categories."
}
