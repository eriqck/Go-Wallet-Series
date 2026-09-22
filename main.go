package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// deposit method
func (w *wallet) Deposit(amount float64) error {

	if amount <= 0 {
		return ErrInvalidAmount
	}

	w.balance += amount
	w.AddTransaction("Deposit", amount)
	return nil

}

// withdraw method
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

// AddTransaction method
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
func (w *wallet) showTransactions() {
	fmt.Println()
	fmt.Println("=====Transaction History====")

	if len(w.transactions) == 0 {
		fmt.Println("No transaction yet!")
	} else {
		for i, t := range w.transactions {
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

//NewWallet() that creates wallets

func NewWallet(balance float64, currency, accountNumber, owner string) *wallet {
	return &wallet{
		balance:       balance,
		currency:      currency,
		accountNumber: accountNumber,
		owner:         owner,
	}

}

func main() {
	fmt.Println("=========================================")
	fmt.Println("       Welcome to the Go Wallet!")
	fmt.Println("=========================================")

	reader := bufio.NewReader(os.Stdin)

	wallets := map[string]*wallet{
		"1030942": NewWallet(10000, "KES", "1030942", "Penelope"),
		"1030943": NewWallet(5000, "USD", "1030943", "Girshom"),
		"1030944": NewWallet(7500, "KES", "1030944", "Eric"),
	}

	fmt.Println("Enter your account number:")

	accountNumber, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Account number not found!")
		return
	}

	accountNumber = strings.TrimSpace(accountNumber)

	currentWallet, exists := wallets[accountNumber]

	if !exists {
		fmt.Println("Account not found!")
		return
	}

	running := true

	for running {

		showMenu()

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Failed to read input")
			continue
		}

		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please enter a valid choice")
			continue
		}
		switch choice {
		case 1:
			fmt.Print("Enter deposit amount: ")

			input, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("Failed to read deposit Amount")
				continue
			}
			input = strings.TrimSpace(input)
			deposit, err := strconv.ParseFloat(input, 64)

			if err != nil {
				fmt.Println("Please enter a valid amount")
				continue
			}

			//recored or reject a deposit
			err = currentWallet.Deposit(deposit)

			if err != nil {
				if errors.Is(err, ErrInvalidAmount) {
					fmt.Println("Deposit Amount must be greater than zero!")
				} else {
					fmt.Println("Deposit failed!", err)
				}
			} else {
				fmt.Println("Deposit successful! New balance:",
					currentWallet.currency,
					currentWallet.balance)
			}

		case 2:
			fmt.Print("Enter withdrawal amount: ")
			input, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("Failed to read withdrawal amount!")
				continue
			}

			input = strings.TrimSpace(input)

			withdraw, err := strconv.ParseFloat(input, 64)

			if err != nil {
				fmt.Println("Please enter a valid amount")
				continue
			}

			err = currentWallet.Withdraw(withdraw)

			if err != nil {
				if errors.Is(err, ErrInsufficientFunds) {
					fmt.Println("Insufficient Funds!")
				} else if errors.Is(err, ErrInvalidAmount) {
					fmt.Println("Withdrawal failed!", err)
				}
			} else {
				fmt.Println("Withdrawal successful! New balance:",
					currentWallet.currency,
					currentWallet.balance)
			}

		case 3:
			showBalance(currentWallet)

		case 4:
			currentWallet.showTransactions()

		case 5:
			fmt.Println("Goodbye!")
			running = false

		default:
			fmt.Println("Please enter a valid menu option")

		}

	}

}
