package cmd

type AnalyseCommand struct {
}

func (a AnalyseCommand) Execute([]string) {
	//TODO implement me
	panic("implement me")
}

func (a AnalyseCommand) Description() string {
	return "Analyzes the type, description and code of the transactions in the CSV files"
}
