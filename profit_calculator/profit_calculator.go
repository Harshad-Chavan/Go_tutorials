package main

import "fmt"

func main() {
	//declare required data
	var revenue float64
	var expenses float64
	var taxRate float64

	//fetch data from the user
	fmt.Print("Enter revevnur in $ : ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expense in $ : ")
	fmt.Scan(&expenses)

	fmt.Print("Enter Tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses

	eat := ebt * (1 - taxRate/100)

	ratio := ebt / eat

	fmt.Println(ebt)
	fmt.Println(eat)
	fmt.Println(ratio)

}
