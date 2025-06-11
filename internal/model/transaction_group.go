package model

const income = "Income"
const expense = "Expense"

type TransactionGroups struct {
	Map map[string]string
}

var TransactionGroupMappings = &TransactionGroups{
	Map: map[string]string{
		"Direct Credit":     income,
		"Visa Refund":       income,
		"Automatic Payment": income,
		"Transfer":          income,
		"Payment":           income,
		"Bank Fee":          expense,
		"Debit Interest":    expense,
		"Atm Debit":         expense,
		"ATM Cash Deposit":  income,
		"Eft-Pos":           expense,
		"Bill Payment":      expense,
		"Deposit":           income,
		"Loan Payment":      expense,
		"EFTPOS":            expense,
		"Visa Purchase":     expense,
		"Direct Debit":      expense,
	},
}

type TransactionGroup struct {
	Name                 string
	Type                 string
	Amount               float64
	NumberOfTransactions int
	Transactions         []*Transaction
}
