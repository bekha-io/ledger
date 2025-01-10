package types

type AccountType string

const (
	AccountTypeAssets      = "assets"
	AccountTypeLiabilities = "liabilities"
	AccountTypeEquities    = "equity"
	AccountTypeIncome      = "income"
	AccountTypeExpenses    = "expenses"
)

var AccountTypes = []AccountType{AccountTypeAssets, AccountTypeIncome, AccountTypeLiabilities, AccountTypeEquities, AccountTypeExpenses}
