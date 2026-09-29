
package main

import (
	"errors"
	"time"
)

var (
	ErrInvalidAmount     = errors.New("invalid amount")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type wallet struct {
	balance       float64
	currency      string
	accountNumber string
	owner         string
	transactions  []transaction
}

type transaction struct {
	transactionType string
	amount          float64
	timestamp       time.Time
}

func NewWallet(balance float64, currency, accountNumber, owner string) *wallet {
	return &wallet{
		balance:       balance,
		currency:      currency,
		accountNumber: accountNumber,
		owner:         owner,
	}
}

func (w *wallet) Deposit(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	w.balance += amount
	w.AddTransaction("Deposit", amount)

	return nil
}

func (w *wallet) Withdraw(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	w.AddTransaction("Withdrawal", amount)

	return nil
}

func (w *wallet) AddTransaction(transactionType string, amount float64) {
	w.transactions = append(w.transactions, transaction{
		transactionType: transactionType,
		amount:          amount,
		timestamp:       time.Now(),
	})
}