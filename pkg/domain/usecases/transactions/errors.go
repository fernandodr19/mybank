package transactions

import "errors"

var (
	ErrAccountNotFound     = errors.New("account not found")
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrInsufficientCredit  = errors.New("insufficient credit")
)
