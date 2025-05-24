package internal

import (
	"bankr/internal/model"
	"fmt"
)

type Processor interface {
	Process(transactions []*model.Transaction) error
}

type TransactionProcessor struct {
}

func (p *TransactionProcessor) Process(transactions []*model.Transaction) error {
	fmt.Println("=> Processing transactions...")
	transactionsByCategory := model.GroupBy(transactions, func(t *model.Transaction) string {
		return t.Type
	})
	var categorySummary *model.Category
	for category, transactions := range transactionsByCategory {
		categorySummary = &model.Category{
			Name: category,
			Amount: model.Sum(transactions, func(t *model.Transaction) float64 {
					return t.Amount
				}),
				NumberOfTransactions: len(transactions),
				//Transactions: transactions,
			}
		PrettyPrintJson(categorySummary)
	}
	return nil
}