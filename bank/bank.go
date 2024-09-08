package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to read from file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to read from file")
	}
	return balance, nil

}

func writeBalanceToFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceTxt), 0644)

}

func main() {

	var accountBalance float64
	var err error
	accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR ! ")
		fmt.Println(err)
		fmt.Println("--------------------------------")
	}

	fmt.Println("Welcome to the bank")

	// for loop without any condition runs like a while loop
	// for i := 0; i < 2; i++
	for {

		fmt.Println("What do you want to do with you account ?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("your choice : ")
		fmt.Scan(&choice)

		//same flow using swtich statement
		switch choice {
		case 1:
			fmt.Println("Balance :", accountBalance)
		case 2:
			var amount float64
			fmt.Println("how much you want to deposit :")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Invalic amount must be greater than zero")
				continue
			} else {
				accountBalance += amount
				fmt.Println("Updated balance amount : ", accountBalance)
				writeBalanceToFile(accountBalance)
			}
		case 3:
			var amount float64
			fmt.Println("how much you want to Withdraw :")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Invalic amount must be greater than zero")
				continue
			}
			if amount > accountBalance {
				fmt.Println("Amount cant be greater than accountBalance")
				continue
			}
			accountBalance -= amount
			fmt.Println("Updated balance amount : ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Exiting")
			fmt.Println("Thank you !! for banking with us")
			return

		}

		// 	if choice == 1 {
		// 		fmt.Println("Balance :", accountBalance)
		// 	} else if choice == 2 {
		// 		var amount float64
		// 		fmt.Println("how much you want to deposit :")
		// 		fmt.Scan(&amount)
		// 		if amount <= 0 {
		// 			fmt.Println("Invalic amount must be greater than zero")
		// 			continue
		// 		} else {
		// 			accountBalance += amount
		// 			fmt.Println("Updated balance amount : ", accountBalance)
		// 		}

		// 	} else if choice == 3 {
		// 		var amount float64
		// 		fmt.Println("how much you want to Withdraw :")
		// 		fmt.Scan(&amount)
		// 		if amount <= 0 {
		// 			fmt.Println("Invalic amount must be greater than zero")
		// 			continue
		// 		}
		// 		if amount > accountBalance {
		// 			fmt.Println("Amount cant be greater than accountBalance")
		// 			continue
		// 		}
		// 		accountBalance -= amount
		// 		fmt.Println("Updated balance amount : ", accountBalance)
		// 	} else {
		// 		fmt.Println("Exiting")
		// 		break
		// 	}
		// }

		// fmt.Println("Thank you !! for banking with us")

	}
}
