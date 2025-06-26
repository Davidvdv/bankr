package internal

import (
	"bankr/internal/model"
	"fmt"
)

type Processor interface {
	Process(transactions []*model.Transaction) []*model.TransactionGroup
}

type TransactionProcessor struct {
}

func (p *TransactionProcessor) Process(transactions []*model.Transaction) []*model.TransactionGroup {
	fmt.Println("=> Processing transactions...")
	groupedTransactions := model.GroupBy(transactions, func(t *model.Transaction) string {
		return t.Type
	})
	transactionGroups := make([]*model.TransactionGroup, 0, len(groupedTransactions))
	for group, transactions := range groupedTransactions {
		transactionGroup := &model.TransactionGroup{
			Name: group,
			Type: model.TransactionGroupMappings.Map[group],
			Amount: model.Sum(model.Map(transactions, func(t *model.Transaction) float64 {
				return t.Amount
			})),
			NumberOfTransactions: len(transactions),
			Transactions:         transactions,
		}
		transactionGroups = append(transactionGroups, transactionGroup)
	}
	fmt.Printf("=> Successfully processed %d transactions\n", len(transactions))
	return transactionGroups
}
