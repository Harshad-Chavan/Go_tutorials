package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func outputText(text string) {
	fmt.Print(text)

}

// basic use case (preferred way)
func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}

// another way of returning data
// func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
// 	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv = fv / math.Pow(1+inflationRate/100, years)
// 	return
// }

// basic program to calculate return on investment
func main() {

	//variable declaration
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	//var investmentAmount, years, expectedReturnRate float64 = 1000, 10, 5.5

	//investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	// getting values from the user and assignment
	outputText("Investment amount : ")
	fmt.Scan(&investmentAmount)

	outputText("years : ")
	fmt.Scan(&years)

	outputText("expected return rate  : ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	//fmt.Println("Future Value: ", futureValue)
	//fmt.Println("Furture values adjusted with inflation:", futureRealValue)

	formattedFV := fmt.Sprintf("Formatted Future value %.2f", futureValue)
	fmt.Println(formattedFV)

	formattedFRV := fmt.Sprintf("Formatted Future value %.2f", futureRealValue)
	fmt.Println(formattedFRV)

}
