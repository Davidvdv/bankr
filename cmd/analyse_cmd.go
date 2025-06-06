package cmd

type AnalyseCommand struct {
}

func (a *AnalyseCommand) Execute([]string) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (a *AnalyseCommand) Description() string {
	return "Analyses the type, description and code of the transactions in the CSV files"
}
