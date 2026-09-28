package main

import (
	"errors"
	"testing"
)

func TestDeposit(t *testing.T) {
	w := NewWallet(1000, "KES", "1030945", "Eric")

	err := w.Deposit(500)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if w.balance != 1500 {
		t.Errorf("expected balance to be 1500, got %v", w.balance)
	}
}

func TestDepositInvalidAmount(t *testing.T) {
	tests := []struct {
		name   string
		amount float64
	}{
		{
			name:   "zero amount",
			amount: 0,
		},
		{
			name:   "negative amount",
			amount: -100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w := NewWallet(1000, "KES", "1030943", "Eric")
			err := w.Deposit(tt.amount)

			if !errors.Is(err, ErrInvalidAmount) {
				t.Fatalf("Expected ErrInvalidAmount, got %v", err)
			}

			if w.balance != 1000 {
				t.Errorf("Expected balance to remain 1000, got %.2f", w.balance)
			}
		})
	}
}

//function to test invalid withdrawals

func TestInvalidWithdrawalAmount(t *testing.T) {

	tests := []struct {
		name   string
		amount float64
	}{
		{
			name:   "zero amount",
			amount: 0,
		},

		{
			name:   "negative amount",
			amount: -100,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			w := NewWallet(1000, "KES", "1030943", "Eric")
			err := w.Withdraw(tt.amount)

			if !errors.Is(err, ErrInvalidAmount) {

				t.Errorf("expected ErrInvalidAmount, got %v", err)

			}

			if w.balance != 1000 {

				t.Errorf("expected balance to remain 1000, got %.2f", w.balance)
			}

		})

	}
}

func TestWithdrawInsufficientFunds(t *testing.T) {

	tests := []struct {
		name   string
		amount float64
	}{
		{
			name:   "withdraw greater than balance",
			amount: 1500,
		},
		{
			name:   "withdraw far greater than balance",
			amount: 15000,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			w := NewWallet(1000, "KES", "1030943", "Eric")
			err := w.Withdraw(tt.amount)

			if !errors.Is(err, ErrInsufficientFunds) {
				t.Fatalf("Expected ErrInsufficientFunds, got %v", err)
			}

			if w.balance != 1000 {
				t.Errorf("Expected balance to remain 1000, got %.2f", w.balance)
			}
		})

	}

}

func TestWithdraw(t *testing.T) {
	w := NewWallet(2000, "KES", "1030955", "Eric")
	err := w.Withdraw(700)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if w.balance != 1300 {
		t.Errorf("Expected balance to be 1300, got %.2f", w.balance)
	}

}

func TestDepositRecordTransaction(t *testing.T) {
	w := NewWallet(1000, "KES", "1030945", "Eric")
	err := w.Deposit(500)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(w.transactions) != 1 {
		t.Fatalf("Expected 1 transaction, got %d", len(w.transactions))
	}

	transaction := w.transactions[0]

	if transaction.transactionType != "Deposit" {
		t.Errorf("expected a deposit type transaction, got %s", transaction.transactionType)
	}

	if transaction.amount != 500 {
		t.Errorf("expected transaction amount to be 1500, got %.2f", transaction.amount)
	}
}

func TestWithdrawRecordsTransactions(t *testing.T) {
	w := NewWallet(1000, "KES", "1030943", "Eric")
	err := w.Withdraw(300)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(w.transactions) != 1 {
		t.Fatalf("expected 1 transaction, got %d", len(w.transactions))
	}

	transaction := w.transactions[0]

	if transaction.transactionType != "Withdrawal" {
		t.Errorf("expected transaction type to be Withdrawal, got %s", transaction.transactionType)
	}

	if transaction.amount != 300 {
		t.Errorf("expected transaction amount to be 300, got %.2f", transaction.amount)
	}
}

func TestFailedWithdrawalDoesNotRecordTransaction(t *testing.T) {
	w := NewWallet(1000, "KSH", "1030943", "Eric")

	err := w.Withdraw(1500)
	if !errors.Is(err, ErrInsufficientFunds) {
		t.Fatalf("expected ErrInsufficientFunds, got %v", err)
	}

	if len(w.transactions) != 0 {
		t.Errorf("expected no transactions, got %d", len(w.transactions))
	}

	if w.balance != 1000 {
		t.Errorf("expected wallet balance to remain 1000, got %.2f", w.balance)
	}

}
