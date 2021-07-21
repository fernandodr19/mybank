package operations

type Operation int

var (
	Debit      Operation = 1
	Credit     Operation = 2
	Withdrawal Operation = 3
	Payment    Operation = 4
)
