package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance float64
	var err error
	accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR ! ")
		fmt.Println(err)
		fmt.Println("--------------------------------")
		// panic("Cannot continue")
	}

	fmt.Println("Welcome to the bank")
	fmt.Println("Reach us at --> ", randomdata.PhoneNumber())
	presentOptions()

	// for loop without any condition runs like a while loop
	// for i := 0; i < 2; i++
	for {

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
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
