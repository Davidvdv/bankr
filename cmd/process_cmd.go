package cmd

type ProcessCommand struct{}

func (p ProcessCommand) Execute([]string) {
	//TODO implement me
	panic("implement me")
}

func (p ProcessCommand) Description() string {
	return "Process the transactions in the CSV files and generate a summary of the transactions and their categories."
}
