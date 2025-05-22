package model

import (
	"errors"
	"strconv"
	"time"
	"fmt"
)

// Type,Details,Particulars,Code,Reference,Amount,Date,ForeignCurrencyAmount,ConversionCharge
type Transaction struct {
	Type                  string
	Details               string
	Particulars           string
	Code                  string
	Reference             string
	Amount                float64
	Date                  time.Time
	ForeignCurrencyAmount string
	ConversionCharge      string
}

func createTransaction(entry []string) (*Transaction, error) {
	if len(entry) < 9 {
		return nil, errors.New("invalid transaction line")
	}
	amount, _ := strconv.ParseFloat(entry[5], 64)
	date, _ := time.Parse("02/01/2006", entry[6])
	return &Transaction{
		Type:                  entry[0],
		Details:               entry[1],
		Particulars:           entry[2],
		Code:                  entry[3],
		Reference:             entry[4],
		Amount:                amount,
		Date:                  date,
		ForeignCurrencyAmount: entry[7],
		ConversionCharge:      entry[8],
	}, nil
}

func BuildTransactions(entries [][]string) []*Transaction {
		transactions := make([]*Transaction, 0, len(entries))
	for _, entry := range entries {
		transaction, err := createTransaction(entry)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		transactions = append(transactions, transaction)
	}
	return transactions
}