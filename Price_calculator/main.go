package main

import (
	"example/price_calculator/prices"
	"fmt"
)

func main() {

	// var prices []float64 = []float64{10, 20, 30, 40}
	var taxRates []float64 = []float64{0, 0.7, 0.1, 0.5}

	var result map[float64][]float64 = make(map[float64][]float64)

	for _, taxrate := range taxRates {
		PriceJob := prices.NewTaxIncludedPriceJob(taxrate)
		PriceJob.Process()

	}

	fmt.Print(result)
}
