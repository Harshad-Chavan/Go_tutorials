package main

import (
	"example/price_calculator/filemanager"
	"example/price_calculator/prices"
	"fmt"
)

func main() {

	// var prices []float64 = []float64{10, 20, 30, 40}

	//price job objects will run for these tax rates
	var taxRates []float64 = []float64{0, 0.7, 0.1, 0.5}

	doneChans := make([](chan bool), len(taxRates))

	for index, taxrate := range taxRates {
		IoManager := filemanager.New("prices/prices.txt", fmt.Sprintf("result_%.0f_prices.json", taxrate*100))
		//IoManager := cmdmanager.New()
		PriceJob := prices.NewTaxIncludedPriceJob(IoManager, taxrate)
		doneChans[index] = make(chan bool)
		go PriceJob.Process(doneChans[index])

	}

	for _, donechan := range doneChans {
		<-donechan
	}

	fmt.Print("Program Completed")
}
