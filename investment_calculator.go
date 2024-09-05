package main

import (
	"fmt"
	"math"
)

// basic program to calculate return on investment
func main() {

	//var investmentAmount float64 = 1000
	//var years float64 = 10
	//var expectedReturnRate float64 = 5.5

	//var investmentAmount, years, expectedReturnRate float64 = 1000, 10, 5.5

	const inflationRate = 2.5
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
