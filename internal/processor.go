package internal

import (
	"bankr/internal/model"
	"fmt"
)

type Processor interface {
	Process(transactions []*model.Transaction)
}

type TransactionProcessor struct {
}

func (p *TransactionProcessor) Process(transactions []*model.Transaction) {
	fmt.Println("=> Processing transactions...")
	transactionsByCategory := model.GroupBy(transactions, func(t *model.Transaction) string {
		return t.Type
	})
	for category, transactions := range transactionsByCategory {
		categorySummary := &model.Category{
			Name: category,
			Type: model.Categories.CategoryByType[category],
			Amount: model.Sum(model.Map(transactions, func(t *model.Transaction) float64 {
				return t.Amount
			})),
			NumberOfTransactions: len(transactions),
			//Transactions: transactions,
		}
		PrettyPrintJson(categorySummary)
	}
}
