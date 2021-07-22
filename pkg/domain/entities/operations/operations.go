package operations

type Operation int

var (
	Debit      Operation = 1
	Credit     Operation = 2
	Withdrawal Operation = 3
	Payment    Operation = 4
)

func (o Operation) String() string {
	switch o {
	case Debit:
		return "debit"
	case Credit:
		return "credit"
	case Withdrawal:
		return "withdrawal"
	case Payment:
		return "payment"
	}

	return ""
}
