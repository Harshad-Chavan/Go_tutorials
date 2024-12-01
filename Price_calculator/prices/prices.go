package prices

import (
	"example/price_calculator/conversion"
	"example/price_calculator/filemanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

// contructor for the price job obj
func NewTaxIncludedPriceJob(taxrate float64) *TaxIncludedPriceJob {
	// returns a taxincluded price job

	newStruc := TaxIncludedPriceJob{
		TaxRate:     taxrate,
		InputPrices: []float64{},
	}
	//sending a pointer
	return &newStruc

}

// Load prices from a file one price on each line
func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := filemanager.ReadLines("prices/prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices

}

// takes pointer to the job and deos not make a copy
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	job.TaxIncludedPrices = result
	filemanager.WriteJson(fmt.Sprintf("result_%.0f_prices.json", job.TaxRate*100), job)

}
