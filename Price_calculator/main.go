package main

import "fmt"

func main() {

	var prices []float64 = []float64{10, 20, 30, 40}
	var taxRates []float64 = []float64{0, 0.7, 0.1, 0.5}

	var result map[float64][]float64 = make(map[float64][]float64)

	for _, taxrate := range taxRates {
		var taxIncludedPrices []float64 = make([]float64, len(prices))
		for index, price := range prices {
			taxIncludedPrices[index] = calculateTax(taxrate, price)
		}
		result[taxrate] = taxIncludedPrices

	}

	fmt.Print(result)
}

func calculateTax(taxrate float64, price float64) float64 {
	return price * (1 + taxrate)
}
