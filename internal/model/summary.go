package model

import "time"

type Summary struct {
	TotalAmountSpent    float64
	TotalAmountReceived float64
	NumberOfTransactions int
	StartDate time.Time
	EndDate   time.Time
}

func BuildSummary(transactions []*Transaction) *Summary {
	sum := func (transactions []*Transaction) float64 {
		total := 0.0
		for _, t := range transactions {
			total += t.Amount
		}
		return total
	}
	return &Summary{
		TotalAmountSpent: sum(transactions),
		TotalAmountReceived: func () float64 {
			debitTransactions := make([]*Transaction, 0, len(transactions))
			for _, t := range transactions {
				if t.Amount > 0 {
					debitTransactions = append(debitTransactions, t)
				}
			}
			return sum(debitTransactions)
		}(),
		NumberOfTransactions: len(transactions),
		StartDate: transactions[len(transactions)-1].Date,
		EndDate: transactions[0].Date,
	}
}
