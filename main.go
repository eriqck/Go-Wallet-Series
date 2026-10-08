package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	walletpkg "wallet/internal/wallet"
)

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
func showBalance(w *walletpkg.Wallet) {
	fmt.Println("Current Balance:",
		w.Currency(),
		w.Balance(),
	)
}

// show transactions
func showTransactions(w *walletpkg.Wallet) {
	fmt.Println()
	fmt.Println("===== Transaction History =====")

	transactions := w.Transactions()

	if len(transactions) == 0 {
		fmt.Println("No transaction yet!")
		return
	}

	for i, t := range transactions {
		fmt.Printf(
			"%d. %s: %s %d at %s\n",
			i+1,
			t.TransactionType,
			w.Currency(),
			t.Amount,
			t.Timestamp.Format("2006-01-02 15:04:05"),
		)
	}
}

func main() {
	fmt.Println("=========================================")
	fmt.Println("       Welcome to the Go Wallet!")
	fmt.Println("=========================================")

	reader := bufio.NewReader(os.Stdin)

	wallets := map[string]*walletpkg.Wallet{
		"1030942": walletpkg.NewWallet(10000, "KES", "1030942", "Penelope"),
		"1030943": walletpkg.NewWallet(5000, "USD", "1030943", "Girshom"),
		"1030944": walletpkg.NewWallet(7500, "KES", "1030944", "Eric"),
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
			deposit, err := strconv.ParseInt(input, 10, 64)

			if err != nil {
				fmt.Println("Please enter a valid amount")
				continue
			}

			//recored or reject a deposit
			err = currentWallet.Deposit(deposit)

			if err != nil {
				if errors.Is(err, walletpkg.ErrInvalidAmount) {
					fmt.Println("Deposit Amount must be greater than zero!")
				} else {
					fmt.Println("Deposit failed!", err)
				}
			} else {
				fmt.Println("Deposit successful! New balance:",
					currentWallet.Currency(),
					currentWallet.Balance())
			}

		case 2:
			fmt.Print("Enter withdrawal amount: ")
			input, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("Failed to read withdrawal amount!")
				continue
			}

			input = strings.TrimSpace(input)

			withdraw, err := strconv.ParseInt(input, 10, 64)

			if err != nil {
				fmt.Println("Please enter a valid amount")
				continue
			}

			err = currentWallet.Withdraw(withdraw)

			if err != nil {
				if errors.Is(err, walletpkg.ErrInsufficientFunds) {
					fmt.Println("Insufficient Funds!")
				} else if errors.Is(err, walletpkg.ErrInvalidAmount) {
					fmt.Println("Withdrawal failed!", err)
				}
			} else {
				fmt.Println("Withdrawal successful! New balance:",
					currentWallet.Currency(),
					currentWallet.Balance())
			}

		case 3:
			showBalance(currentWallet)

		case 4:
			showTransactions(currentWallet)

		case 5:
			fmt.Println("Goodbye!")
			running = false

		default:
			fmt.Println("Please enter a valid menu option")

		}

	}

}
