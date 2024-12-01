package main

import (
	"example/price_calculator/prices"
	"fmt"
)

func main() {

	// var prices []float64 = []float64{10, 20, 30, 40}

	//price job objects will run for these tax rates
	var taxRates []float64 = []float64{0, 0.7, 0.1, 0.5}

	for _, taxrate := range taxRates {
		PriceJob := prices.NewTaxIncludedPriceJob(taxrate)
		PriceJob.Process()

	}

	fmt.Print("Program Completed")
}
