package wallet

import (
	"errors"
	"time"
)

var (
	ErrInvalidAmount     = errors.New("invalid amount")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type Wallet struct {
	balance       float64
	currency      string
	accountNumber string
	owner         string
	transactions  []Transaction
}

type Transaction struct {
	TransactionType string
	Amount          float64
	Timestamp       time.Time
}

func NewWallet(balance float64, currency, accountNumber, owner string) *Wallet {
	return &Wallet{
		balance:       balance,
		currency:      currency,
		accountNumber: accountNumber,
		owner:         owner,
	}
}

func (w *Wallet) Currency() string {
	return w.currency
}

func (w *Wallet) Balance() float64 {
	return w.balance
}

func (w *Wallet) AccountNumber() string {
	return w.accountNumber
}

func (w *Wallet) Owner() string {
	return w.owner
}

func (w *Wallet) Transactions() []Transaction {
	return w.transactions
}

func (w *Wallet) Deposit(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	w.balance += amount
	w.AddTransaction("Deposit", amount)

	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
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

func (w *Wallet) AddTransaction(transactionType string, amount float64) {
	w.transactions = append(w.transactions, Transaction{
		TransactionType: transactionType,
		Amount:          amount,
		Timestamp:       time.Now(),
	})
}
