package transactions

import "errors"

var (
	ErrAccountNotFound     = errors.New("account not found")
	ErrInvalidAccID        = errors.New("invalid account")
	ErrInvalidAmount       = errors.New("invalid amount")
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrInsufficientCredit  = errors.New("insufficient credit")
)
