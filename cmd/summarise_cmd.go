package cmd

type SummariseCommand struct {
}

func (s SummariseCommand) Execute([]string) {
	//TODO implement me
	panic("implement me")
}

func (s SummariseCommand) Description() string {
	return "Provides a summary of the transactions in the CSV files"
}
