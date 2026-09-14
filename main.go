package main

import (
	"errors"
	"fmt"
	"time"
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

// deposit method
func (w *wallet) Deposit(amount float64) error {

	if amount <= 0 {
		return errors.New("invalid deposit amount")
	}

	w.balance += amount
	w.AddTransaction("Deposit", amount)
	return nil

}

// withdraw method
func (w *wallet) Withdraw(amount float64) error {

	if amount <= 0 {
		return errors.New("invalid withdrawal amount")
	}
	if amount > w.balance {
		return errors.New("insufficient funds")
	}

	w.balance -= amount
	w.AddTransaction("Withdrawal", amount)
	return nil
}

//AddTransaction method
func (w *wallet) AddTransaction(transactionType string, amount float64) {
	w.transactions = append(w.transactions, transaction{
		transactionType: transactionType,
		amount:          amount,
		timestamp:       time.Now(),
	})
}

/*
func findWallet(accountNumber string, wallets []wallet) *wallet {
	for i := range wallets {
		if wallets[i].accountNumber == accountNumber {
			return &wallets[i]
		}
	}
	return nil
}
*/

// Display Menu
func showMenu() {
	fmt.Println()
	fmt.Println("1.Deposit")
	fmt.Println("2.Withdraw")
	fmt.Println("3.Check Balance")
	fmt.Println("4.Transaction History")
	fmt.Println("5.Exit")

}

// Show Balance
func showBalance(w *wallet) {
	fmt.Println("Current Balance:",
		w.currency,
		w.balance,
	)
}

// show transactions
func showTransactions(transactions []transaction, w *wallet) {
	fmt.Println()
	fmt.Println("=====Transaction History====")

	if len(transactions) == 0 {
		fmt.Println("No transaction yet!")
	} else {
		for i, t := range transactions {
			fmt.Printf("%d. %s: %s %.2f at %s\n",
				i+1,
				t.transactionType,
				w.currency,
				t.amount,
				t.timestamp.Format("2006-01-02 15:04:05"),
			)

		}
	}
}

func main() {
	fmt.Println("=========================================")
	fmt.Println("       Welcome to the Go Wallet!")
	fmt.Println("=========================================")

	wallets := map[string]*wallet{
		"1030942": {
			balance:       10000,
			currency:      "KES",
			accountNumber: "1030942",
			owner:         "Penelope",
		},

		"1030943": {
			balance:       5000,
			currency:      "USD",
			accountNumber: "1030943",
			owner:         "Girshom",
		},

		"1030944": {
			balance:       7500,
			currency:      "KES",
			accountNumber: "1030944",
			owner:         "Eric",
		},
	}

	fmt.Println("Enter your account number:")
	var accountNumber string
	fmt.Scan(&accountNumber)

	currentWallet, exists := wallets[accountNumber]

	if !exists {
		fmt.Println("Account not found!")
		return
	}

	running := true

	for running {

		showMenu()

		choice := 0
		fmt.Print("Enter a choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter deposit amount: ")
			var deposit float64
			fmt.Scan(&deposit)

			//recored or reject a deposit
			err := currentWallet.Deposit(deposit)

			if err != nil {
				fmt.Println("Deposit failed!", err)
			} else {
				fmt.Println("Deposit successful! New balance:",
					currentWallet.currency,
					currentWallet.balance)
			}

		case 2:
			fmt.Print("Enter withdrawal amount: ")
			var withdraw float64
			fmt.Scan(&withdraw)

			err := currentWallet.Withdraw(withdraw)

			if err != nil {
				fmt.Println("Withdrawal failed!", err)
			} else {
				fmt.Println("Withdrawal successful! New balance:",
					currentWallet.currency,
					currentWallet.balance)
			}

		case 3:
			showBalance(currentWallet)

		case 4:
			showTransactions(currentWallet.transactions, currentWallet)

		case 5:
			fmt.Println("Goodbye!")
			running = false

		}

	}

}
