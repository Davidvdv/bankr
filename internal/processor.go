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
	groupedTransactions := model.GroupBy(transactions, func(t *model.Transaction) string {
		return t.Type
	})
	for group, transactions := range groupedTransactions {
		transactionGroup := &model.TransactionGroup{
			Name: group,
			Type: model.TransactionGroupMappings.Map[group],
			Amount: model.Sum(model.Map(transactions, func(t *model.Transaction) float64 {
				return t.Amount
			})),
			NumberOfTransactions: len(transactions),
			//Transactions: transactions,
		}
		PrettyPrintJson(transactionGroup)
	}
}
