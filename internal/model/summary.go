package model

import "time"

type Summary struct {
	NumberOfAccounts			 int
	TotalAmountSpent    float64
	TotalAmountReceived float64
	NumberOfTransactions int
	StartDate time.Time
	EndDate   time.Time
}

type Category struct {
	Name string
	Amount float64
	NumberOfTransactions int
	Transactions []*Transaction
}

func BuildSummary(transactions []*Transaction, numberOfAccounts int) *Summary {
	sum := func (transactions []*Transaction) float64 {
		total := 0.0
		for _, t := range transactions {
			total += t.Amount
		}
		return total
	}
	return &Summary{
		NumberOfAccounts:     numberOfAccounts,
		TotalAmountSpent: sum(transactions),
		TotalAmountReceived: sum(Filter(transactions, func(t *Transaction) bool {
			return t.Amount > 0
		})),
		NumberOfTransactions: len(transactions),
		StartDate: transactions[len(transactions)-1].Date,
		EndDate: transactions[0].Date,
	}
}
