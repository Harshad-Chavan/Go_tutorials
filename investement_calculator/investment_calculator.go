package main

import (
	"fmt"
	"math"
)

// basic program to calculate return on investment
func main() {

	//variable declaration
	var investmentAmount float64

	var years float64
	var expectedReturnRate float64

	//var investmentAmount, years, expectedReturnRate float64 = 1000, 10, 5.5

	const inflationRate = 2.5
	//investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	// getting values from the user and assignment
	fmt.Print("Investment amount : ")
	fmt.Scan(&investmentAmount)

	fmt.Print("years : ")
	fmt.Scan(&years)

	fmt.Print("expected return rate  : ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
