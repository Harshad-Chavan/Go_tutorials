package main

import "fmt"

func main() {
	//declare required data
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	//fetch data from the user
	revenue := getUserInput("Revenue :")
	expenses := getUserInput("Expenses :")
	taxRate := getUserInput("Tax Rate :")

	ebt := revenue - expenses

	eat := ebt * (1 - taxRate/100)

	ratio := ebt / eat

	fmt.Println(ebt)
	fmt.Println(eat)
	fmt.Println(ratio)

}

func getUserInput(infotext string) float64 {
	var userInput float64
	fmt.Print(infotext)
	fmt.Scan(&userInput)
	return userInput
}
