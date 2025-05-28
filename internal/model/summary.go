package model

import "time"

type Summary struct {
	NumberOfAccounts     int
	TotalAmountSpent     float64
	TotalAmountReceived  float64
	NumberOfTransactions int
	StartDate            time.Time
	EndDate              time.Time
}

func BuildSummary(transactions []*Transaction, numberOfAccounts int) *Summary {
	return &Summary{
		NumberOfAccounts: numberOfAccounts,
		TotalAmountSpent: Sum(Map(Filter(transactions, func(t *Transaction) bool {
			return t.Amount < 0
		}), func(t *Transaction) float64 {
			return t.Amount
		})),
		TotalAmountReceived: Sum(Map(Filter(transactions, func(t *Transaction) bool {
			return t.Amount > 0
		}), func(t *Transaction) float64 {
			return t.Amount
		})),
		NumberOfTransactions: len(transactions),
		StartDate:            transactions[len(transactions)-1].Date,
		EndDate:              transactions[0].Date,
	}
}
